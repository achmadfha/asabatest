<template>
  <div class="h-full login-bg bg-slate-50">
    <div class="flex h-full justify-center items-center">
      <div class="h-max min-w-[16rem] w-1/4 max-w-[24rem] text-center items-center">
        <div class="inline-flex mt-4 mb-8 items-center">
          <img src="@/assets/weave.png" class="h-12 mr-2" />
          <h1 class="font-bold text-4xl font-mono">Weaver</h1>
        </div>

        <div v-if="showLogin">
            <el-form ref="loginFormRef" :model="loginUser" size="large" :rules="rules" show-message>
              <el-form-item prop="name">
                <el-input v-model="loginUser.email" placeholder="admin">
                  <template #prefix>
                    <User />
                  </template>
                </el-input>
              </el-form-item>

              <el-form-item prop="password">
                <el-input v-model="loginUser.password" type="password" placeholder="123456" show-password>
                  <template #prefix>
                    <Lock />
                  </template>
                </el-input>
              </el-form-item>
            </el-form>

            <el-button class="w-full" type="primary" size="large" @click="login(loginFormRef)">SIGN IN</el-button>
            <div class="w-full flex mt-[0.25rem]">
              <div class="w-full text-right">
                <el-button link @click="showLogin=false">SIGN UP</el-button>
              </div>
            </div>

        </div>

          <div v-if="showLogin == false">
            <el-form ref="registerFormRef" :model="registerUser" label-position="top" :rules="rules"
              label-width="auto" size="large">
              <el-form-item label="Email" prop="email">
                <el-input placeholder="email" v-model="registerUser.email"></el-input>
              </el-form-item>
              <el-form-item label="Password" prop="password">
                <el-input placeholder="password" type="password" minlength="6" v-model="registerUser.password"></el-input>
              </el-form-item>
            </el-form>

            <el-button class="w-full" type="primary" size="large" @click="register(registerFormRef)">SIGN UP</el-button>
            <div class="mt-[0.25rem] text-right">
              <el-button link @click="showLogin=true">SIGN IN</el-button>
            </div>
          </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-bg {
  background-image: url('@/assets/login-bg.svg');
  background-repeat: no-repeat;
  background-size: 100% auto;
  background-position: 0 100%;
}
</style>

<script setup>
import { ElMessage, ElNotification } from "element-plus"
import { User, Lock } from '@icon-park/vue-next'
import { ref, reactive } from 'vue'
import request from '@/axios'
import { useRouter } from 'vue-router'

const router = useRouter();

const loginFormRef = ref();
const registerFormRef = ref();
const activeTab = ref('login');

const showLogin = ref(true);

const loginUser = reactive({
  email: "",
  password: "",
});
const registerUser = reactive({
  email: "",
  password: "",
});
const rules = reactive({
  password: [
    { required: true, message: 'Please input password', trigger: 'blur' },
    { min: 6, message: 'Length should be great than 6', trigger: 'blur' }
  ],
  email: [
    { required: true, message: 'Please input email', trigger: 'blur' },
    { type: 'email', message: 'Please input correct email address', trigger: ['blur', 'change'] },
  ]
});

const login = async (form) => {
  if (!form) {
    return;
  }

  let email = loginUser.email;

  let success = function () {
    ElNotification.success({
      title: 'Login Success',
      message: 'Hi~ ' + email,
      showClose: true,
      duration: 1500,
    });
    router.push('/');
  };

  await form.validate((valid, fields) => {
    if (valid) {
      request.post("/api/v1/auth/login", {
        email: loginUser.email,
        password: loginUser.password,
        setCookie: true,
      }).then((response) => {
        if (response.data.responseCode === "2000100") {;
          document.cookie = `access_token=${response.data.data.access_token}; path=/`;
          success();
        } else {;
          ElMessage({
            message: response.data.responseMessage,
            type: "error",
          });
        }
      }).catch((error) => {
        console.error("Login error:", error);
        ElMessage({
          message: "An error occurred while logging in.",
          type: "error",
        });
      });
    } else {
      console.log("Input invalid", fields);
      ElMessage({
        message: "Input invalid" + fields,
        type: "error",
      });
    }
  });
};


const register = async (form) => {
  if (!form) {
    return
  }

  await form.validate((valid, fields) => {
    if (valid) {
      request.post("/api/v1/auth/register", {
        email: registerUser.email,
        password: registerUser.password
      }).then((response) => {
        ElMessage({
          message: 'Register success',
          type: 'success',
        })
        loginUser.email = registerUser.email;
        loginUser.password = registerUser.password;
        activeTab.value = 'login';
      })
    } else {
      ElMessage({
        message: "Input invalid",
        type: "error",
      });
    }
  });
};
</script>
