<script setup>
import { ref, onMounted } from 'vue';

    const name = ref("Rattapon");
    const status = ref("active");
    const tasks = ref(["Task one", "Task two", "Task three"]);

    const toggleStatus = () => {
      status.value = status.value === "active" ? "inactive" : "active";
    };
    const newTask = ref("");

    const addTask = () => {
      if (newTask.value.trim() !== "") {
        tasks.value.push(newTask.value.trim());
        newTask.value = ""; // Clear the input field after adding the task
      }
    };

onMounted(async () => {
  try {
    const respone = await fetch('https://jsonplaceholder.typicode.com/todos');
    const data = await respone.json();
    tasks.value = data.map((task) => task.title)
  } catch (error) {
    console.log('Error fetch tasks')
  }
});
</script>

<template>
  <h1>{{ name }}</h1>
  <p v-if="status === 'active'">User status: Active</p>
  <p v-else>User status: Inactive</p>

  <form @submit="addTask">
    <label for="newTask">Add Task</label>
    <input type="text" id="newTask" v-model="newTask" required>
    <button type="submit">Submit</button>
  </form>

  <h3>Tasks:</h3>
  <ul>
    <li v-for="(task, index) in tasks" :key="index">
      <span>{{ task }}</span>
      <button @click="tasks.splice(index, 1)">x</button>
    </li>
  </ul>
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


</style>
