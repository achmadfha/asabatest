<template>
  <div class="w-full flex justify-center">
    <div class="flex flex-col w-full h-full px-[4rem] py-[2rem] space-y-[1rem]">
      <el-dialog v-model="showCreate" top="5vh" title="Add Transactions" width="50%">
        <el-form ref="createFormRef" label-position="top" label-width="auto">
          <div v-for="(transaction, index) in newTransactions" :key="index">
            <el-form-item label="Code" prop="code" required>
              <el-input v-model="transaction.items_code" maxlength="10"/>
              <span class="text-gray-400">Max 10 characters</span>
            </el-form-item>
            <el-form-item label="Amount" prop="amount" required>
              <el-input-number v-model="transaction.quantity" controls-position="right" :min="0"/>
              <span class="text-gray-400 ms-4">Quantity</span>
            </el-form-item>
            <el-form-item label="Transaction Type" prop="transaction_type" required>
              <el-select v-model="transaction.transaction_type" placeholder="Select Transaction Type">
                <el-option label="IN" value="IN"/>
                <el-option label="OUT" value="OUT"/>
              </el-select>
              <span class="text-gray-400 ms-4">Transaction Type</span>
            </el-form-item>
          </div>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button type="primary" @click="createTransactions">Confirm</el-button>
            <el-button @click="addAnotherTransaction">+ Add More</el-button>
            <el-button @click="showCreate = false; resetForm()">Cancel</el-button>
          </span>
        </template>
      </el-dialog>

      <el-dialog v-model="showEditDialog" top="5vh" title="Update Transaction" width="50%">
        <template v-for="(transaction, index) in additionalTransactions" :key="index">
          <el-form :ref="`updateTransactionFormRef${index}`" :model="transaction" label-position="top" label-width="auto">
            <el-form-item label="Code" prop="items_code">
              <el-input v-model="transaction.items_code" disabled/>
            </el-form-item>
            <el-form-item label="Amount" prop="quantity" required>
              <el-input-number v-model="transaction.quantity" controls-position="right" :min="0"/>
              <span class="text-gray-400 ms-4">Quantity</span>
            </el-form-item>
            <el-form-item label="Transaction Type" prop="transaction_type" required>
              <el-select v-model="transaction.transaction_type" placeholder="Select Transaction Type">
                <el-option label="IN" value="IN"/>
                <el-option label="OUT" value="OUT"/>
              </el-select>
              <span class="text-gray-400 ms-4">Transaction Type</span>
            </el-form-item>
          </el-form>
        </template>
        <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="updateTransactions">Confirm</el-button>
          <el-button @click="showEditDialog = false">Cancel</el-button>
        </span>
        </template>
      </el-dialog>

      <div class="flex overflow-hidden rounded-md shadow-md border">
        <div class="flex w-full h-[5rem] items-center">
          <History class="ml-[1rem]" theme="filled" size="42" fill="#94A3B8"/>
          <span class="m-[0.75rem] text-2xl font-600">Transactions</span>
        </div>
      </div>

      <el-card class="w-full h-max">
        <template #header>
          <div class="flex w-full space-x-[2rem]">
            <el-select class="w-1/3" v-model="selectedTransactionType" filterable placeholder="Select Transaction Type"
                       @change="filterItems">
              <el-option label="All" value=""/>
              <el-option label="IN" value="IN"/>
              <el-option label="OUT" value="OUT"/>
            </el-select>

            <el-input v-model="search" placeholder="Type to search">
              <template #prefix>
                <el-icon>
                  <Search/>
                </el-icon>
              </template>
            </el-input>

            <el-button type="primary" plain :icon="History" @click="showCreate = true">Create</el-button>
            <el-button type="primary" plain :icon="Edit" @click="editSelectedItems">Edit Selected Transactions</el-button>
          </div>
        </template>
        <el-table :data="paginatedTransactions" class="w-full max-h-full" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="50">

          </el-table-column>
          <el-table-column prop="items_code" label="Code"/>
          <el-table-column prop="quantity" label="Quantity"/>
          <el-table-column prop="transaction_type" label="Transaction Type">
            <template #default="scope">
              <el-tag :type="scope.row.transaction_type === 'IN' ? 'success' : 'danger'">
                {{ scope.row.transaction_type }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="created_at" label="Created At" :formatter="formatDate"/>
          <el-table-column prop="updated_at" label="Updated At" :formatter="formatDate"/>
          <el-table-column label="Operation" min-width="150px">
            <template #default="scope">
              <el-popover :visible="showDelete == scope.$index" placement="top" :width="180">
                <template #reference>
                  <el-button size="small" type="danger" @click="showDelete = scope.$index" :icon="Delete" circle
                             class="wl-[1rem]"/>
                </template>
                <p>Are you sure to delete this item?</p>
                <div class="my-[0.5rem]">
                  <el-button size="small" text @click="showDelete = -1">cancel</el-button>
                  <el-button size="small" type="danger" @click="deleteTransaction(scope.row)">confirm</el-button>
                </div>
              </el-popover>
            </template>
          </el-table-column>
        </el-table>

        <el-pagination
            v-model:currentPage="currentPage"
            :page-size="pageSize"
            :total="totalTransactions"
            layout="prev, pager, next"
            @current-change="handlePageChange"
        />

      </el-card>
    </div>
  </div>
</template>

<script setup>
import {Edit, Delete, History, Search} from '@icon-park/vue-next';
import {ref, computed, watch} from 'vue';
import axios from "axios";
import {ElMessage} from "element-plus";
import {onMounted} from 'vue';
import {getAccessToken} from '@/utils';
import moment from "moment/moment";

const transactions = ref([]);
const allTransactions = ref([]);
const showDelete = ref(-1);
const currentPage = ref(1);
const pageSize = ref(10);
const totalTransactions = computed(() => transactions.value.length);
const selectedTransactionType = ref('All');
const search = ref('');
const showCreate = ref(false);
const showEditDialog = ref(false);
const additionalTransactions = ref([]);
const selectedItems = ref([]);

const editSelectedItems = () => {
  if (selectedItems.value.length === 0) {
    ElMessage.warning('Please select items to edit.');
    return;
  }

  additionalTransactions.value = selectedItems.value.map(item => ({ ...item }));
  showEditDialog.value = true;
};


const handleSelectionChange = (selected) => {
  selectedItems.value = selected;
};


const newTransactions = ref([{
  items_code: '',
  quantity: 0,
  transaction_type: ''
}]);

const resetForm = () => {
  newTransactions.value = [{
    items_code: '',
    quantity: 0,
    transaction_type: ''
  }];
};

const addAnotherTransaction = () => {
  newTransactions.value.push({
    items_code: '',
    quantity: 0,
    transaction_type: ''
  });
};

const fetchTransactions = async (transactionType) => {
  try {
    const accessToken = getAccessToken();
    if (!accessToken) {
      throw new Error('Access token not found.');
    }

    let url = '/api/v1/transaction';
    if (transactionType && transactionType !== 'All') {
      url += `?transactionsType=${transactionType}`;
    }

    const response = await axios.get(url, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    transactions.value = response.data.data;
    allTransactions.value = response.data.data;
  } catch (error) {
    ElMessage.error('Error fetching transactions. Please try again.');
  }
};

onMounted(fetchTransactions);

const paginatedTransactions = computed(() => {
  const startIndex = (currentPage.value - 1) * pageSize.value;
  const endIndex = startIndex + pageSize.value;
  return transactions.value.slice(startIndex, endIndex);
});

const handlePageChange = (page) => {
  currentPage.value = page;
};

watch(selectedTransactionType, (newType) => {
  fetchTransactions(newType);
});

const formatDate = (dateString) => {
  return moment(dateString).format('DD-MM-YYYY');
};

watch(search, (newSearchTerm) => {
  const lowerCaseSearchTerm = newSearchTerm.toLowerCase();
  transactions.value = allTransactions.value.filter(transaction =>
      transaction && (transaction.items_code.toLowerCase().includes(lowerCaseSearchTerm) ||
          transaction.quantity.toString().includes(lowerCaseSearchTerm) ||
          transaction.transaction_type.toLowerCase().includes(lowerCaseSearchTerm))
  );
});

const createTransactions = async () => {
  try {
    const accessToken = getAccessToken();
    if (!accessToken) {
      throw new Error('Access token not found.');
    }

    for (const transaction of newTransactions.value) {
      const response = await axios.post('/api/v1/transaction', transaction, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });

      if (response.data.responseCode === '2000300') {
        ElMessage({
          message: 'Success Create Transaction',
          type: 'success',
        });
      } else {
        ElMessage({
          message: response.data.responseMessage,
          type: 'error',
        });
      }
    }

    fetchTransactions();
    showCreate.value = false;
    resetForm();
  } catch (error) {
    if (error.response && error.response.data.responseCode === '4030303') {
      ElMessage({
        message: error.response.data.responseMessage,
        type: 'error',
      });
    } else {
      ElMessage({
        message: 'An error occurred while creating the transaction.',
        type: 'error',
      });
    }
    fetchTransactions();
    showCreate.value = false;
    resetForm();
  }
};

const deleteTransaction = async (row) => {
  try {
    const accessToken = getAccessToken();
    if (!accessToken) {
      throw new Error('Access token not found.');
    }

    const response = await axios.delete(`/api/v1/transaction/${row.transaction_id}`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    if (response.data.responseCode === '2000300') {
      ElMessage({
        message: 'Success Delete Transaction',
        type: 'success',
      });
      fetchTransactions();
    } else {
      ElMessage({
        message: response.data.responseMessage,
        type: 'error',
      });
    }
  } catch (e) {
    ElMessage({
      message: 'An error occurred while deleting transaction.',
      type: 'error',
    });
  }
};

</script>