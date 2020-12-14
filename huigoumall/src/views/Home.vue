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
import Header from '@/components/Header.vue'
import Aside from '@/components/Aside.vue'

export default {
  name: 'Home',
  components: {
    Header,
    Aside
  }
}
</script>
