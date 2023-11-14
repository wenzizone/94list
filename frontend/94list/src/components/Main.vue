<template>
  <a-layout>
    <a-layout style="margin: 16px 0; padding: 0 24px 24px">
      <h2>前台解析中心 | 94list-vue-go</h2>
      <a-layout-content
          :style="{ background: '#fff', padding: '24px', margin: 0, minHeight: '280px' }"
        >
        <a-form
          name="basic"
          :label-col="{ span: 8 }"
          :wrapper-col="{ span: 16 }"
          autocomplete="off"
          style="max-width: 600px"
        >
          <a-form-item
            label="链接"
            name="链接"
            :rules="[{ required: true, message: '请填入百度网盘分享链接' }]"
          >
            <a-input />
          </a-form-item>

          <a-form-item
            label="密码"
            name="密码"
          >
            <a-input />
          </a-form-item>

          <a-form-item
            label="当前路径"
            name="当前路径"
          >
            <a-input value="/" disabled />
          </a-form-item>

          <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
            <a-button type="primary" html-type="submit">Submit</a-button>
          </a-form-item>
        </a-form>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

