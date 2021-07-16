<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="系统名称">
          <el-input
            placeholder="搜索条件"
            v-model="searchInfo.asset_system_name"
          />
        </el-form-item>
        <el-form-item label="产品经理">
          <el-input
            placeholder="搜索条件"
            v-model="searchInfo.asset_system_manager"
          />
        </el-form-item>
        <el-form-item label="域名">
          <el-input
            placeholder="搜索条件"
            v-model="searchInfo.asset_system_domain"
          />
        </el-form-item>
        <el-form-item label="外网ip">
          <el-input placeholder="搜索条件" v-model="searchInfo.extranet_ip" />
        </el-form-item>
        <el-form-item label="外网端口">
          <el-input placeholder="搜索条件" v-model="searchInfo.extranet_port" />
        </el-form-item>
        <el-form-item label="内网ip">
          <el-input placeholder="搜索条件" v-model="searchInfo.intranet_ip" />
        </el-form-item>
        <el-form-item label="内网端口">
          <el-input placeholder="搜索条件" v-model="searchInfo.intranet_port" />
        </el-form-item>
        <el-form-item label="归属测试环境" prop="is_test_environment">
          <el-select
            v-model="searchInfo.is_test_environment"
            clear
            placeholder="请选择"
          >
            <el-option key="true" label="是" value="true"> </el-option>
            <el-option key="false" label="否" value="false"> </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="web状态码">
          <el-input
            placeholder="搜索条件"
            v-model="searchInfo.web_status_code"
          />
        </el-form-item>
        <el-form-item label="重点资产" prop="is_important_asset">
          <el-select
            v-model="searchInfo.is_important_asset"
            clear
            placeholder="请选择"
          >
            <el-option key="true" label="是" value="true"> </el-option>
            <el-option key="false" label="否" value="false"> </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input placeholder="搜索条件" v-model="searchInfo.more_record" />
        </el-form-item>
        <el-form-item label="url链接">
          <el-input placeholder="搜索条件" v-model="searchInfo.url" />
        </el-form-item>
        <el-form-item label="web指纹信息">
          <el-input placeholder="搜索条件" v-model="searchInfo.fingerprint" />
        </el-form-item>

        <el-form-item>
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-search"
            @click="onSubmit"
            >查询</el-button
          >
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-plus"
            @click="openDialog"
            >新增</el-button
          >
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false"
                >取消</el-button
              >
              <el-button size="mini" type="primary" @click="onDelete"
                >确定</el-button
              >
            </div>
            <el-button
              slot="reference"
              icon="el-icon-delete"
              size="mini"
              type="danger"
              style="margin-left: 10px"
              >批量删除</el-button
            >
          </el-popover>
        </el-form-item>

        <el-form-item label="" prop="export">
          <el-button
            type="success"
            icon="el-icon-download"
            size="mini"
            @click="exportResult"
          >
            {{ export_button_name }}
          </el-button>
        </el-form-item>
        <el-form-item label="" prop="export">
          <el-upload
            type="success"
            class="excel-btn"
            :action="`${path}/asset_manage_system/import_excel`"
            size="mini"
            :on-success="loadExcel"
            :headers="{ 'x-token': token }"
            :show-file-list="false"
            
          >
            <el-button size="small" type="primary" icon="el-icon-upload2"
              >导入</el-button
            >
          </el-upload>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{
          scope.row.UpdatedAt | formatDate
        }}</template>
      </el-table-column>
      <el-table-column label="系统名称" prop="asset_system_name" width="120" />
      <el-table-column
        label="产品经理"
        prop="asset_system_manager"
        width="120"
      />
      <el-table-column label="域名" prop="asset_system_domain" width="180" />
      <el-table-column label="外网ip" prop="extranet_ip" width="150" />
      <el-table-column label="外网端口" prop="extranet_port" width="100" />
      <el-table-column label="内网ip" prop="intranet_ip" width="100" />
      <el-table-column label="内网端口" prop="intranet_port" width="100" />
      <el-table-column
        label="归属测试环境"
        prop="is_test_environment"
        width="50"
      >
        <template slot-scope="scope">{{
          scope.row.is_test_environment | formatBoolean
        }}</template>
      </el-table-column>
      <el-table-column label="web状态码" prop="web_status_code" width="80" />
      <el-table-column label="web截图" prop="web_screenshot" width="120" />
      <el-table-column label="重点资产" prop="is_important_asset" width="80">
        <template slot-scope="scope">{{
          scope.row.is_important_asset | formatBoolean
        }}</template>
      </el-table-column>
      <el-table-column label="url链接" prop="url" width="120" />
      <el-table-column label="web指纹信息" prop="fingerprint" width="120" />
      <el-table-column label="备注" prop="more_record" width="150" />
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button
            size="small"
            type="primary"
            icon="el-icon-edit"
            class="table-button"
            @click="updateAsset_manage_system(scope.row)"
            >变更</el-button
          >
          <el-button
            type="danger"
            icon="el-icon-delete"
            size="mini"
            @click="deleteRow(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <el-dialog
      :before-close="closeDialog"
      :visible.sync="dialogFormVisible"
      title="弹窗操作"
    >
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="系统名称:">
          <el-input
            v-model="formData.asset_system_name"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="产品经理:">
          <el-input
            v-model="formData.asset_system_manager"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="域名:">
          <el-input
            v-model="formData.asset_system_domain"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="外网ip:">
          <el-input
            v-model="formData.extranet_ip"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="外网端口:">
          <el-input
            v-model="formData.extranet_port"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="内网ip:">
          <el-input
            v-model="formData.intranet_ip"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="内网端口:">
          <el-input
            v-model="formData.intranet_port"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="归属测试环境:">
          <el-switch
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            v-model="formData.is_test_environment"
            clearable
          ></el-switch>
        </el-form-item>
        <el-form-item label="web状态码:">
          <el-input
            v-model.number="formData.web_status_code"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="web截图:">
          <el-input
            v-model="formData.web_screenshot"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="重点资产:">
          <el-switch
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            v-model="formData.is_important_asset"
            clearable
          ></el-switch>
        </el-form-item>
        <el-form-item label="url链接:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="web指纹信息:">
          <el-input
            v-model="formData.fingerprint"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input
            v-model="formData.more_record"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import {
  createAsset_manage_system,
  deleteAsset_manage_system,
  deleteAsset_manage_systemByIds,
  updateAsset_manage_system,
  findAsset_manage_system,
  getAsset_manage_systemList,
  exportAsset_manage_system_resultsByIds,
  exportAsset_manage_system_resultsByConditions,
  loadExcelData,
} from "@/api/asset_manage_system"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "Asset_manage_system",
  mixins: [infoList],
  computed: {
    ...mapGetters("user", ["userInfo", "token"]),
  },
  data() {
    return {
      listApi: getAsset_manage_systemList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      export_button_name: "导出所有筛选结果",
      path: path,
      formData: {
        asset_system_name: "",
        asset_system_manager: "",
        asset_system_domain: "",
        extranet_ip: "",
        extranet_port: "",
        intranet_ip: "",
        intranet_port: "",
        is_test_environment: false,
        web_status_code: 0,
        web_screenshot: "",
        is_important_asset: false,
        more_record: "",
        url: "",
        fingerprint: "",
      },
    };
  },
  filters: {
    formatDate: function (time) {
      if (time !== null && time !== "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    },
  },
  async created() {
    await this.getTableData();
  },
  methods: {
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 100;
      if (this.searchInfo.is_test_environment === "") {
        this.searchInfo.is_test_environment = null;
      }
      if (this.searchInfo.is_important_asset === "") {
        this.searchInfo.is_important_asset = null;
      }
      this.getTableData();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
      if (this.multipleSelection.length !== 0) {
        this.export_button_name = "导出当前选中结果";
      } else {
        this.export_button_name = "导出所有筛选结果";
      }
    },
    deleteRow(row) {
      this.$confirm("确定要删除吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        this.deleteAsset_manage_system(row);
      });
    },
    async onDelete() {
      const ids = [];
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: "warning",
          message: "请选择要删除的数据",
        });
        return;
      }
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          ids.push(item.ID);
        });
      // console.log(ids)
      const res = await deleteAsset_manage_systemByIds({ ids });
      if (res.code === 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--;
        }
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    loadExcel() {
      this.listApi = loadExcelData
      this.getTableData()
    },
    exportResult() {
      const filename = "ExportResult.xlsx";
      if (this.multipleSelection.length === 0) {
        console.log("导出所有结果中");
        exportAsset_manage_system_resultsByConditions(
          this.searchInfo,
          filename
        );
      } else {
        // console.log(this.tableData);
        //导出选中筛选条件的结果
        const iterms = [];
        this.multipleSelection &&
          this.multipleSelection.map((item) => {
            // console.log(item);
            iterms.push(item);
          });

        exportAsset_manage_system_resultsByIds(iterms, filename);
        // console.log(iterms)
        // if (res.code === 0) {
        //   alert("导出成功");
        // }
      }
    },
    async updateAsset_manage_system(row) {
      const res = await findAsset_manage_system({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.reasset_manage_system;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        asset_system_name: "",
        asset_system_manager: "",
        asset_system_domain: "",
        extranet_ip: "",
        extranet_port: 0,
        intranet_ip: "",
        intranet_port: 0,
        is_test_environment: false,
        web_status_code: 0,
        web_screenshot: "",
        is_important_asset: false,
        more_record: "",
        url: "",
        fingerprint: "",
      };
    },
    async deleteAsset_manage_system(row) {
      const res = await deleteAsset_manage_system({ ID: row.ID });
      if (res.code === 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAsset_manage_system(this.formData);
          break;
        case "update":
          res = await updateAsset_manage_system(this.formData);
          break;
        default:
          res = await createAsset_manage_system(this.formData);
          break;
      }
      if (res.code === 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功",
        });
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
  },
};
</script>

<style>
</style>
