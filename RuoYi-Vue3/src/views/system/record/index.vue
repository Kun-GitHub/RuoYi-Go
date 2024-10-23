<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" size="small" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="分类" prop="classify">
        <el-input
          v-model="queryParams.classify"
          placeholder="请输入分类"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="标题" prop="title">
        <el-input
          v-model="queryParams.title"
          placeholder="请输入标题"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="发布时间" prop="timeShow">
        <el-input
          v-model="queryParams.timeShow"
          placeholder="请输入发布时间"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="信息类型" prop="stageName">
        <el-input
          v-model="queryParams.stageName"
          placeholder="请输入信息类型"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="来源平台" prop="platformName">
        <el-input
          v-model="queryParams.platformName"
          placeholder="请输入来源平台"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="业务类型" prop="classifyShow">
        <el-input
          v-model="queryParams.classifyShow"
          placeholder="请输入业务类型"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="行业" prop="tradeShow">
        <el-input
          v-model="queryParams.tradeShow"
          placeholder="请输入行业"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="省份" prop="districtShow">
        <el-input
          v-model="queryParams.districtShow"
          placeholder="请输入省份"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="${comment}" prop="url">
        <el-input
          v-model="queryParams.url"
          placeholder="请输入${comment}"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="信息类型" prop="stageShow">
        <el-input
          v-model="queryParams.stageShow"
          placeholder="请输入信息类型"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="${comment}" prop="titleShow">
        <el-input
          v-model="queryParams.titleShow"
          placeholder="请输入${comment}"
          clearable
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          type="primary"
          plain
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
          v-hasPermi="['system:record:add']"
        >新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="success"
          plain
          icon="el-icon-edit"
          size="mini"
          :disabled="single"
          @click="handleUpdate"
          v-hasPermi="['system:record:edit']"
        >修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="danger"
          plain
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
          v-hasPermi="['system:record:remove']"
        >删除</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="warning"
          plain
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
          v-hasPermi="['system:record:export']"
        >导出</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="recordList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="主键" align="center" prop="id" />
      <el-table-column label="分类" align="center" prop="classify" />
      <el-table-column label="标题" align="center" prop="title" />
      <el-table-column label="发布时间" align="center" prop="timeShow" />
      <el-table-column label="信息类型" align="center" prop="stageName" />
      <el-table-column label="来源平台" align="center" prop="platformName" />
      <el-table-column label="业务类型" align="center" prop="classifyShow" />
      <el-table-column label="行业" align="center" prop="tradeShow" />
      <el-table-column label="省份" align="center" prop="districtShow" />
      <el-table-column label="${comment}" align="center" prop="url" />
      <el-table-column label="信息类型" align="center" prop="stageShow" />
      <el-table-column label="${comment}" align="center" prop="titleShow" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['system:record:edit']"
          >修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
            v-hasPermi="['system:record:remove']"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryParams.pageNum"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />

    <!-- 添加或修改【请填写功能名称】对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="500px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="分类" prop="classify">
          <el-input v-model="form.classify" placeholder="请输入分类" />
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入标题" />
        </el-form-item>
        <el-form-item label="发布时间" prop="timeShow">
          <el-input v-model="form.timeShow" placeholder="请输入发布时间" />
        </el-form-item>
        <el-form-item label="信息类型" prop="stageName">
          <el-input v-model="form.stageName" placeholder="请输入信息类型" />
        </el-form-item>
        <el-form-item label="来源平台" prop="platformName">
          <el-input v-model="form.platformName" placeholder="请输入来源平台" />
        </el-form-item>
        <el-form-item label="业务类型" prop="classifyShow">
          <el-input v-model="form.classifyShow" placeholder="请输入业务类型" />
        </el-form-item>
        <el-form-item label="行业" prop="tradeShow">
          <el-input v-model="form.tradeShow" placeholder="请输入行业" />
        </el-form-item>
        <el-form-item label="省份" prop="districtShow">
          <el-input v-model="form.districtShow" placeholder="请输入省份" />
        </el-form-item>
        <el-form-item label="${comment}" prop="url">
          <el-input v-model="form.url" placeholder="请输入${comment}" />
        </el-form-item>
        <el-form-item label="信息类型" prop="stageShow">
          <el-input v-model="form.stageShow" placeholder="请输入信息类型" />
        </el-form-item>
        <el-form-item label="${comment}" prop="titleShow">
          <el-input v-model="form.titleShow" placeholder="请输入${comment}" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listRecord, getRecord, delRecord, addRecord, updateRecord } from "@/api/system/record";

export default {
  name: "Record",
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 显示搜索条件
      showSearch: true,
      // 总条数
      total: 0,
      // 【请填写功能名称】表格数据
      recordList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        classify: null,
        title: null,
        timeShow: null,
        stageName: null,
        platformName: null,
        classifyShow: null,
        tradeShow: null,
        districtShow: null,
        url: null,
        stageShow: null,
        titleShow: null,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        title: [
          { required: true, message: "标题不能为空", trigger: "blur" }
        ],
        url: [
          { required: true, message: "$comment不能为空", trigger: "blur" }
        ],
      }
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询【请填写功能名称】列表 */
    getList() {
      this.loading = true;
      listRecord(this.queryParams).then(response => {
        this.recordList = response.rows;
        this.total = response.total;
        this.loading = false;
      });
    },
    // 取消按钮
    cancel() {
      this.open = false;
      this.reset();
    },
    // 表单重置
    reset() {
      this.form = {
        id: null,
        classify: null,
        title: null,
        timeShow: null,
        stageName: null,
        platformName: null,
        classifyShow: null,
        tradeShow: null,
        districtShow: null,
        url: null,
        stageShow: null,
        titleShow: null,
        createTime: null
      };
      this.resetForm("form");
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageNum = 1;
      this.getList();
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.resetForm("queryForm");
      this.handleQuery();
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length!==1
      this.multiple = !selection.length
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();
      this.open = true;
      this.title = "添加【请填写功能名称】";
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getRecord(id).then(response => {
        this.form = response.data;
        this.open = true;
        this.title = "修改【请填写功能名称】";
      });
    },
    /** 提交按钮 */
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.id != null) {
            updateRecord(this.form).then(response => {
              this.$modal.msgSuccess("修改成功");
              this.open = false;
              this.getList();
            });
          } else {
            addRecord(this.form).then(response => {
              this.$modal.msgSuccess("新增成功");
              this.open = false;
              this.getList();
            });
          }
        }
      });
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const ids = row.id || this.ids;
      this.$modal.confirm('是否确认删除【请填写功能名称】编号为"' + ids + '"的数据项？').then(function() {
        return delRecord(ids);
      }).then(() => {
        this.getList();
        this.$modal.msgSuccess("删除成功");
      }).catch(() => {});
    },
    /** 导出按钮操作 */
    handleExport() {
      this.download('system/record/export', {
        ...this.queryParams
      }, `record_${new Date().getTime()}.xlsx`)
    }
  }
};
</script>
