<template>
  <div class="home">
    <Header></Header>
    <el-container>
      <Aside></Aside>
      <el-main>
        <router-view>
        </router-view>
      </el-main>
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
