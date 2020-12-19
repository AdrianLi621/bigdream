<template>
  <div class="home">
    <el-container>
      <Aside></Aside>
      <el-container style="position: relative">
        <Header></Header>
        <el-main>
          <router-view></router-view>
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Header from '@/components/Layout/Header.vue'
import Aside from '@/components/Layout/Aside.vue'

export default {
  name: 'Home',
  components: {
    Header,
    Aside
  }
}
</script>
<style>

</style>
