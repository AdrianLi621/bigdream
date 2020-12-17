<template>
  <div class="goods">
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>商品列表</el-breadcrumb-item>
    </el-breadcrumb>

    <el-row>
      <el-divider content-position="left">搜索条件</el-divider>
      <el-form :inline="true" :model="formInline" class="demo-form-inline" size="mini">
        <el-form-item label="审批人">
          <el-input v-model="formInline.user" placeholder="审批人"></el-input>
        </el-form-item>
        <el-form-item label="活动区域">
          <el-select v-model="formInline.region" placeholder="活动区域">
            <el-option label="区域一" value="shanghai"></el-option>
            <el-option label="区域二" value="beijing"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item>
      </el-form>
      <el-divider></el-divider>
    </el-row>

    <el-table :data="tableData" style="width: 100%">
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column prop="admin_name" label="姓名"></el-table-column>
      <el-table-column prop="admin_password" label="密码"></el-table-column>
      <el-table-column prop="last_login_time" label="最后登录时间" :formatter="dateForma"></el-table-column>
      <el-table-column prop="last_login_ip" label="最后登录IP"></el-table-column>
      <el-table-column prop="created_at" label="注册时间" :formatter="dateForma"></el-table-column>
      <el-table-column prop="updated_at" label="更新时间" :formatter="dateForma"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button type="primary" icon="el-icon-edit" @click="handleClick(scope.row)" size="mini">编辑</el-button>
          <el-button type="danger" icon="el-icon-delete" @click="handleClick(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-row>
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-sizes="[20, 50, 100, 200]"
        :current-page="currentPage"
        :page-size="PageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="totalCount">
      </el-pagination>
    </el-row>
  </div>
</template>

<script>
import moment from "moment"

export default {
  name: 'goods',
  data () {
    return {
      formInline: {
        user: '',
        region: ''
      },
      tableData: [],
      multipleSelection: [],
      totalCount: 1,
      PageSize: 20,
      currentPage: 1
    }
  },
  created () {
    this.getDataList(this.PageSize, this.currentPage)
  },
  methods: {
    onSubmit () {
      console.log('submit!')
    },
    getDataList (pageSize, page) {
      this.$axios.post('adminList', { size: pageSize, page: page },
        {
          // headers:{"Content-Type": "application/x-www-form-urlencoded;charset=utf-8"}
        }).then(res => {
        var response = res.data
        if (response.code !== 1) {
          return this.$message.error(response.msg)
        }
        this.tableData = response.data.list
        this.totalCount = response.data.count
      })
    },
    // 时间格式化转换
    dateForma (row, column) {
      var date = row[column.property]
      if (date === 'undefined') {
        return ''
      }
      return moment(date).format('YYYY-MM-DD HH:mm:ss')
    },
    // 每页显示的条数
    handleSizeChange (val) {
      // 改变每页显示的条数
      this.PageSize = val
      // 注意：在改变每页显示的条数时，要将页码显示到第一页
      this.currentPage = 1
    },
    // 显示第几页
    handleCurrentChange (val) {
      // 改变默认的页数
      this.currentPage = val
    }
  }
}
</script>
