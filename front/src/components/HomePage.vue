<template>
  <div>
    <h1>Calculate Your BMI</h1>

    <!-- Form for user input -->
    <form @submit.prevent="fetchData">
      <label for="height">Height (meters):</label>
      <input
        type="number"
        step="0.01"
        id="height"
        v-model="form.height"
        required
      />
      <br />
      <label for="weight">Weight (kg):</label>
      <input
        type="number"
        step="0.1"
        id="weight"
        v-model="form.weight"
        required
      />
      <br />
      <button type="submit">Calculate BMI</button>
    </form>

    <!-- Display loading status -->
    <div v-if="loading">Loading...</div>

    <!-- Display API response -->
    <div v-else-if="data">
      <h2>Results</h2>
      <pre>{{ data }}</pre>
    </div>

    <!-- Display error message -->
    <div v-if="error" class="error">
      <p>Error: {{ error }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      form: {
        height: null,
        weight: null,
      },
      data: null,
      loading: false,
      error: null,
    };
  },
  methods: {
    async fetchData() {
      // Validate form input before making API call
      if (!this.form.height || !this.form.weight) {
        this.error = 'Please enter valid height and weight.';
        return;
      }

      this.loading = true;
      this.error = null;

      try {
        const response = await axios.get('http://localhost:8080/calculate-bmi', {
          params: {
            height: this.form.height,
            weight: this.form.weight,
          },
        });
        this.data = response.data;
      } catch (err) {
        this.error = err.message || 'An error occurred';
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
.error {
  color: red;
}
form {
  margin-bottom: 20px;
}
button {
  margin-top: 10px;
}
</style>
