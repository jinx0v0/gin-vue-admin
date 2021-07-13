<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="系统名称:">
    <el-input v-model="formData.asset_system_name" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="产品经理:">
    <el-input v-model="formData.asset_system_manager" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="域名:">
    <el-input v-model="formData.asset_system_domain" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="外网ip:">
    <el-input v-model="formData.extranet_ip" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="外网端口:">
    <el-input v-model.number="formData.extranet_port" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="内网ip:">
    <el-input v-model="formData.intranet_ip" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="内网端口:">
    <el-input v-model.number="formData.intranet_port" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="归属测试环境:">
    <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.is_test_environment" clearable ></el-switch>
    </el-form-item>
    
      <el-form-item label="web状态码:">
    <el-input v-model.number="formData.web_status_code" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="web截图:">
    <el-input v-model="formData.web_screenshot" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="重点资产:">
    <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.is_important_asset" clearable ></el-switch>
    </el-form-item>
    
      <el-form-item label="备注:">
    <el-input v-model="formData.more_record" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="url链接:">
    <el-input v-model="formData.url" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="web指纹信息:">
    <el-input v-model="formData.fingerprint" clearable placeholder="请输入" />
    </el-form-item>
    <el-form-item>
        <el-button size="mini" type="primary" @click="save">保存</el-button>
        <el-button size="mini" type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
  createAsset_manage_system,
  updateAsset_manage_system,
  findAsset_manage_system
} from '@/api/asset_manage_system' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Asset_manage_system',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            asset_system_name: '',
            asset_system_manager: '',
            asset_system_domain: '',
            extranet_ip: '',
            extranet_port: 0,
            intranet_ip: '',
            intranet_port: 0,
            is_test_environment: false,
            web_status_code: 0,
            web_screenshot: '',
            is_important_asset: false,
            more_record: '',
            url: '',
            fingerprint: '',
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findAsset_manage_system({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.reasset_manage_system
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createAsset_manage_system(this.formData)
          break
        case 'update':
          res = await updateAsset_manage_system(this.formData)
          break
        default:
          res = await createAsset_manage_system(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>