<template>
  <div class="app-container">
    <h1 class="title">欢迎来到抽奖大转盘</h1>
    
    <div class="content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import router from './router';

export default {
  router,
}
</script>

<style>
.app-container {
  min-height: 90vh;
  padding: 2rem;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  display: flex;
  flex-direction: column;
  align-items: center;
}

.title {
  color: #2c3e50;
  font-size: 2.5rem;
  font-weight: 600;
  margin-bottom: 2rem;
  text-align: center;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
}

.content {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 1rem;
  padding: 2rem;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(5px);
  width: 100%;
  max-width: 800px;
  display: flex;
  justify-content: center;
}

@media (max-width: 768px) {
  .title {
    font-size: 1.8rem;
  }
  
  .content {
    padding: 1rem;
  }
}
</style>