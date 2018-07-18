<template>
  <div class="header-font">
    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            Primary title
          </h1>
          <h2 class="subtitle">
            Primary subtitle
          </h2>
          <h3 class="subtitle">
            <progress class="progress is-small" value="15" max="100">15%</progress>
          </h3>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
  module.exports = {
    data: function () {
      return {
        greeting: 'Hello'
      }
    },
    methods: {
     // async fetchSomething() {
      //  const ip = await this.$axios.$get('http://localhost:8077')
      //  this.ip = ip
      //}
    }
  }
</script>

<style scoped>

</style>
