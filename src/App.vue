<template>
  <div>
    
    <div class="main">
      <component :is="currentComponent" />
    </div>
  </div>
</template>

<script>
import IT09 from './components/it09.vue'
export default {
  components: {IT09},
  data() {
    return {
      current: 1,
    }
  },
  computed: {
    currentComponent() {
     return 'IT09'
    },
  },
}
</script>

<style>
.sidebar {
  float: left;
  width: 120px;
  background: #eee;
  height: 100vh;
  padding: 10px 0;
}
.sidebar button {
  display: block;
  width: 100%;
  margin-bottom: 8px;
  padding: 8px;
  border: none;
  background: #ddd;
  cursor: pointer;
}
.sidebar button.active {
  background: #27ae60;
  color: #fff;
}
.main {
  margin-left: 130px;
  padding: 20px;
}
</style>
