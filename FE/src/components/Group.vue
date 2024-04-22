<template>
  <div class="w-full flex justify-center">
    <div class="flex flex-col w-full h-full px-[4rem] py-[2rem] space-y-[1rem]">
      <el-dialog v-model="showCreate" top="5vh" title="Add Items" width="50%">
        <el-form ref="createFormRef" :model="newItems" label-position="top" label-width="auto">
          <el-form-item label="Code" prop="code" required>
            <el-input v-model="newItems.code" maxlength="10"/>
            <span class="text-gray-400">Max 10 characters</span>
          </el-form-item>
          <el-form-item label="Name" prop="name" required>
            <el-input v-model="newItems.name"/>
            <span class="text-gray-400">Items Name</span>
          </el-form-item>
          <el-form-item label="Amount" prop="amount" required>
            <el-input-number v-model="newItems.amount" controls-position="right" :min="0"/>
            <span class="text-gray-400 ms-4">Items Amount</span>
          </el-form-item>
          <el-form-item label="Description" prop="description" required>
            <el-input v-model="newItems.description" type="textarea"/>
            <span class="text-gray-400">The item describe information</span>
          </el-form-item>
          <el-form-item label="Active" prop="status_active" required>
            <el-switch v-model="newItems.status_active" active-text="True" inactive-text="False"/>
            <span class="text-gray-400 ms-4">Status active (True/False)</span>
          </el-form-item>
        </el-form>
        <template v-if="showAdditionalForms">
          <template v-for="(form, index) in additionalForms" :key="index">
            <el-form ref="additionalFormRef" :model="form" label-position="top" label-width="auto">
              <el-form-item label="Code" prop="code" required>
                <el-input v-model="form.code" maxlength="10"/>
                <span class="text-gray-400">Max 10 characters</span>
              </el-form-item>
              <el-form-item label="Name" prop="name" required>
                <el-input v-model="form.name"/>
                <span class="text-gray-400">Items Name</span>
              </el-form-item>
              <el-form-item label="Amount" prop="amount" required>
                <el-input-number v-model="form.amount" controls-position="right" :min="0"/>
                <span class="text-gray-400 ms-4">Items Amount</span>
              </el-form-item>
              <el-form-item label="Description" prop="description" required>
                <el-input v-model="form.description" type="textarea"/>
                <span class="text-gray-400">The item describe information</span>
              </el-form-item>
              <el-form-item label="Active" prop="status_active" required>
                <el-switch v-model="form.status_active" active-text="True" inactive-text="False"/>
                <span class="text-gray-400 ms-4">Status active (True/False)</span>
              </el-form-item>
            </el-form>
          </template>
        </template>
        <template #footer>
          <span class="dialog-footer">
            <el-button type="primary" @click="createItems">Confirm</el-button>
            <el-button @click="addAnotherItem">+ Add More</el-button>
            <el-button @click="showCreate = false; resetForm()">Cancel</el-button>
          </span>
        </template>
      </el-dialog>

      <el-dialog v-model="showEditDialog" top="5vh" title="Update Item" width="50%">
        <template v-for="(form, index) in additionalForms" :key="index">
          <el-form :ref="`updateFormRef${index}`" :model="form" label-position="top" label-width="auto">
            <el-form-item label="Code" prop="code">
              <el-input v-model="form.code" disabled/>
            </el-form-item>
            <el-form-item label="Name" prop="name">
              <el-input v-model="form.name" disabled/>
            </el-form-item>
            <el-form-item label="Amount" prop="amount" required>
              <el-input-number v-model="form.amount" controls-position="right" :min="0"/>
              <span class="text-gray-400 ms-4">Items Amount</span>
            </el-form-item>
            <el-form-item label="Description" prop="description">
              <el-input v-model="form.description" disabled type="textarea"/>
              <span class="text-gray-400">The item describe information</span>
            </el-form-item>
            <el-form-item label="Active" prop="status_active" required>
              <el-switch v-model="form.status_active" active-text="True" inactive-text="False"/>
              <span class="text-gray-400 ms-4">Status active (True/False)</span>
            </el-form-item>
            <el-form-item label="Transaction Type" prop="transaction_type" required>
              <el-select v-model="form.transaction_type" placeholder="Select Transaction Type">
                <el-option label="IN" value="IN"/>
                <el-option label="OUT" value="OUT"/>
              </el-select>
              <span class="text-gray-400 ms-4">Transaction Type</span>
            </el-form-item>
          </el-form>
        </template>
        <template #footer>
    <span class="dialog-footer">
      <el-button type="primary" @click="updateItems">Confirm</el-button>
      <el-button @click="showEditDialog = false">Cancel</el-button>
    </span>
        </template>
      </el-dialog>

      <div class="flex overflow-hidden rounded-md shadow-md border">
        <div class="flex w-full h-[5rem] items-center">
          <ApplicationTwo class="ml-[1rem]" theme="filled" size="42" fill="#94A3B8"/>
          <span class="m-[0.75rem] text-2xl font-600">Items</span>
        </div>
      </div>

      <el-card class="w-full h-max">
        <template #header>
          <div class="flex w-full space-x-[2rem]">
            <el-input v-model="search" placeholder="Type to search">
              <template #prefix>
                <el-icon>
                  <Search/>
                </el-icon>
              </template>
            </el-input>
            <el-button type="primary" plain :icon="ApplicationTwo" @click="showCreate = true">Create</el-button>
            <el-button type="primary" plain :icon="Edit" @click="editSelectedItems">Edit Selected Items</el-button>
          </div>
        </template>
        <el-table :data="paginatedItems" class="w-full max-h-full" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="50"></el-table-column>
          <el-table-column prop="code" label="Code" width="150"/>
          <el-table-column prop="name" label="Name" width="150"/>
          <el-table-column prop="amount" label="Amount" width="100"/>
          <el-table-column prop="description" label="Description" min-width="200" show-overflow-tooltip/>
          <el-table-column prop="status_active" label="Active" width="100" :formatter="formatActiveStatus"/>
          <el-table-column prop="createdAt" label="Created At" width="150" :formatter="formatDate"/>
          <el-table-column prop="updatedAt" label="Updated At" width="150" :formatter="formatDate"/>
          <el-table-column label="Operation" min-width="120px">
            <template #default="scope">
              <el-button size="small" circle @click="showEditDialog = true; updatedItems = { ...scope.row }"
                         :icon="Edit"/>
              <el-popover :visible="showDelete == scope.$index" placement="top" :width="180">
                <template #reference>
                  <el-button size="small" type="danger" @click="showDelete = scope.$index" :icon="Delete" circle
                             class="wl-[1rem]"/>
                </template>
                <p>Are you sure to delete this item?</p>
                <div class="my-[0.5rem]">
                  <el-button size="small" text @click="showDelete = -1">cancel</el-button>
                  <el-button size="small" type="danger" @click="deleteItem(scope.row)">confirm</el-button>
                </div>
              </el-popover>
            </template>
          </el-table-column>
        </el-table>

        <el-pagination
            v-model:currentPage="currentPage"
            :page-size="pageSize"
            :total="totalItems"
            layout="prev, pager, next"
            @current-change="handlePageChange"
        />

      </el-card>
    </div>
  </div>
</template>

<style scoped>
</style>

<script setup>
import {Edit, Delete, ApplicationTwo, Search} from '@icon-park/vue-next';
import {ref, computed, onMounted, watch} from 'vue';
import {ElMessage} from 'element-plus';
import request from '@/axios';
import {getAccessToken} from '@/utils';
import moment from 'moment';
import axios from "axios";

const items = ref([]);
const showDelete = ref(-1);
const search = ref('');
const currentPage = ref(1);
const pageSize = ref(10);
const totalItems = computed(() => items.value.length);
let allItems = ref([]);
const showCreate = ref(false);
const showEditDialog = ref(false);
const updatedItems = ref({});
const showAdditionalForms = ref(false);
const additionalForms = ref([]);
const selectedItems = ref([]);

const handleSelectionChange = (selected) => {
  selectedItems.value = selected;
};

const editSelectedItems = () => {
  additionalForms.value = [];
  selectedItems.value.forEach(item => {
    additionalForms.value.push({ ...item });
  });

  showEditDialog.value = true;
};

const newItems = ref({
  code: '',
  name: '',
  amount: 0,
  description: '',
  status_active: true,
});

const addAnotherItem = () => {
  console.log('Adding another item...');
  showAdditionalForms.value = true;
  console.log('showAdditionalForms:', showAdditionalForms.value);
  additionalForms.value.push({
    code: '',
    name: '',
    amount: 0,
    description: '',
    status_active: true,
  });
};

const resetForm = () => {
  showAdditionalForms.value = false;
  additionalForms.value = [];
  newItems.value = {
    code: '',
    name: '',
    amount: 0,
    description: '',
    status_active: true,
  };
};

onMounted(async () => {
  fetchItems();
});

const fetchItems = async () => {
  try {
    const accessToken = getAccessToken();
    if (!accessToken) {
      throw new Error('Access token not found.');
    }

    let url = '/api/v1/items';
    const response = await request.get(url, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    items.value = response.data.data;
    allItems.value = response.data.data;
    filterItems();
  } catch (error) {
    ElMessage.error('Error fetching items. Please try again.');
  }
};

const filterItems = () => {
  const lowerCaseSearchTerm = search.value.toLowerCase();
  items.value = allItems.value.filter(item =>
      item && (item.name.toLowerCase().includes(lowerCaseSearchTerm) ||
          item.description.toLowerCase().includes(lowerCaseSearchTerm) ||
          item.code.toLowerCase().includes(lowerCaseSearchTerm))
  );
  currentPage.value = 1;
};

const paginatedItems = computed(() => {
  const startIndex = (currentPage.value - 1) * pageSize.value;
  const endIndex = startIndex + pageSize.value;
  return items.value.slice(startIndex, endIndex);
});

const handlePageChange = (page) => {
  currentPage.value = page;
};

const formatDate = (dateString) => {
  return moment(dateString).format('DD-MM-YYYY');
};

const formatActiveStatus = (row) => {
  return row.status_active ? 'Yes' : 'No';
};

const createItems = async () => {
  try {
    const accessToken = getAccessToken();
    if (!accessToken) {
      throw new Error('Access token not found.');
    }

    const mainFormResponse = await axios.post('/api/v1/items', newItems.value, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    if (mainFormResponse.data.responseCode === '2000200') {
      ElMessage({
        message: 'Success Create Main Items',
        type: 'success',
      });
    }

    const additionalFormResponses = await Promise.all(
        additionalForms.value.map(async (form) => {
          return await axios.post('/api/v1/items', form, {
            headers: {
              Authorization: `Bearer ${accessToken}`,
            },
          });
        })
    );

    additionalFormResponses.forEach((response, index) => {
      if (response.data.responseCode === '2000200') {
        ElMessage({
          message: `Success Create Additional Form ${index + 1}`,
          type: 'success',
        });
      }
    });

    fetchItems();
    showCreate.value = false;
  } catch (e) {
    if (e.response.data.responseCode === '4030203') {
      ElMessage({
        message: e.response.data.responseMessage,
        type: 'error',
      });
    } else {
      ElMessage({
        message: 'An error occurred while creating items.',
        type: 'error',
      });
    }
  }
};

const updateItems = async () => {
  try {
    const accessToken = getAccessToken();
    if (!accessToken) {
      throw new Error('Access token not found.');
    }

    for (const form of additionalForms.value) {
      const response = await axios.put(`/api/v1/items/${form.code}`, form, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });

      if (response.data.responseCode === '2000200') {
        ElMessage({
          message: `Success Update Item ${form.code}`,
          type: 'success',
        });
      } else {
        ElMessage({
          message: response.data.responseMessage,
          type: 'error',
        });
      }
    }

    fetchItems();
    showEditDialog.value = false;
  } catch (e) {
    if (e.response.data.responseCode === '4030203') {
      ElMessage({
        message: e.response.data.responseMessage,
        type: 'error',
      });
    } else {
      ElMessage({
        message: 'An error occurred while updating items.',
        type: 'error',
      });
    }
  }
};

const deleteItem = async (row) => {
  try {
    const accessToken = getAccessToken();
    if (!accessToken) {
      throw new Error('Access token not found.');
    }

    const response = await axios.delete(`/api/v1/items/${row.code}`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    if (response.data.responseCode === '2000200') {
      ElMessage({
        message: 'Success Delete Item',
        type: 'success',
      });
      fetchItems();
    } else {
      ElMessage({
        message: response.data.responseMessage,
        type: 'error',
      });
    }
  } catch (e) {
    ElMessage({
      message: 'An error occurred while deleting item.',
      type: 'error',
    });
  }
};

watch(search, filterItems);
</script>