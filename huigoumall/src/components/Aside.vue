<template>
  <el-aside style="width: 200px;">
    <el-menu class="el-menu-vertical-demo" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b"
             router>
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>权限管理</span>
        </template>
        <el-menu-item index="/admin">管理员列表</el-menu-item>
        <el-menu-item index="/role">角色列表</el-menu-item>
        <el-menu-item index="/permission">权限列表</el-menu-item>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>

</script>

<style scoped>


</style>
