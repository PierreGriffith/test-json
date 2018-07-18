<template>
  <div class="container">
  <section class="container">
    <no-ssr placeholder="Codemirror Loading...">
      <codemirror v-model="code"
                  :options="cmOption"
                  @cursorActivity="onCmCursorActivity"
                  @ready="onCmReady"
                  @focus="onCmFocus"
                  @blur="onCmBlur">
      </codemirror>
    </no-ssr>
  </section>
  </div>
</template>

<script>
  export default {
    data() {
      const code = ` PTDR XD`
      return {
        code,
        cmOption: {
          tabSize: 4,
          foldGutter: true,
          matchBrackets: true,
          styleActiveLine: true,
          lineNumbers: true,
          line: true,
          keyMap: "sublime",
          mode: 'application/json',
          theme: 'blackboard',
        }
      }
    },

    methods: {
      onCmCursorActivity(codemirror) {
        console.log('onCmCursorActivity', codemirror)
      },
      onCmReady(codemirror) {
        console.log('onCmReady', codemirror.setSize(null, 600))
      },
      onCmFocus(codemirror) {
        console.log('onCmFocus', codemirror)
      },
      onCmBlur(codemirror) {
        console.log('onCmBlur', codemirror)
      }
    }
  }
</script>

<style lang="scss" scoped>

  
  .CodeMirror {
    border: 1px solid #eee;
    height: auto;
  }


</style>
