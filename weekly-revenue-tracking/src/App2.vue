<script>
export default {
  data() {
    return {
      name: 'Weekly Revenue Tracking',
      status: 'active',
      tasks: ['Task one', 'Task two', 'Task three'],
      link: 'https://www.youtube.com'
    };
  },

  methods: {
    toggleStatus() {
      if (this.status === 'active') {
        this.status = 'inactive';
      } else {
        this.status = 'active';
      }
    }
  },
};

</script>

<template>
  <h1>{{ name }}</h1>
  <p v-if="status === 'active'">Status: Active</p>
  <p v-else>Status: Inactive</p>

  <h3>Tasks:</h3>
  <ul>
    <li v-for="(task, index) in tasks" :key="index">{{ task }}</li>
  </ul>
  <a v-bind:href="link" target="_blank">Go to YouTube</a><br>

  <!-- <button v-on:click="toggleStatus">Change status</button> -->
  <button @click="toggleStatus">Change status</button>


</template>

<style scoped>
button {
  font-family: 'Segoe UI', 'Helvetica Neue', Arial, 'sans-serif';
  font-size: 1.1rem;
  font-weight: bold;
  background: linear-gradient(90deg, #42b983 0%, #38a1db 100%);
  color: rgb(255, 255, 255);
  border: none;
  border-radius: 12px;
  padding: 12px 24px;
  margin-top: 24px;
  box-shadow: 0 2px 16px rgba(66, 185, 131, 0.08);
  cursor: pointer;
  transition: background 0.2s, color 0.2s, transform 0.2s;
  letter-spacing: 1px;
}
button:hover {
  color: #fff;
  background: green;
  transform: scale(1.04);
}

h1 {
  font-family: 'Segoe UI', 'Helvetica Neue', Arial, 'sans-serif';
  font-size: 2rem;
  font-weight: bold;
  background: linear-gradient(90deg, #42b983 0%, #38a1db 100%);
  color: green;
  -webkit-background-clip: text;
  background-clip: text;
  letter-spacing: 2px;
  border-radius: 12px;
  padding: 16px 16px;
  margin-bottom: 32px;
  box-shadow: 0 2px 16px rgba(66, 185, 131, 0.08);
  transition: transform 0.2s;
  text-align: center;
}
h1:hover {
  color: #fff;
  background: green;
  transform: scale(1.04);
}

</style>
