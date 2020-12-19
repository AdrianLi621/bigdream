<template>
  <el-aside class="aside-layout">
    <el-menu class="el-menu" router>
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>权限管理</span>
        </template>
        <el-menu-item index="/admin">管理员列表</el-menu-item>
        <el-menu-item index="/role">角色列表</el-menu-item>
      </el-submenu>
      <el-submenu index="">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>商品管理</span>
        </template>
        <el-menu-item index="/goods">商品列表</el-menu-item>
        <el-menu-item index="/goods_class">商品分类</el-menu-item>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>

</script>

<style scoped>
.el-aside{
  overflow: hidden;
  flex: 0 0 208px;
  max-width: 208px;
  min-width: 208px;
  width: 208px;
}
.aside-layout{
  position: fixed;
  z-index: 100;
  height: 100%;
  overflow: auto;
  overflow-x: hidden;
  box-shadow: 2px 0 8px 0 rgba(29,35,41,.05);
}
</style>
