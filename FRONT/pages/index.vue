<template>
  <div class="container">

    <h2 class="subtitle has-text-centered">

      <div class="tiny-input">
        <div class="field">
          <div class="control">
            <p class="is-horizontal">
            <input  v-model="TmpServer" class="input is-small is-rounded" type="text" placeholder="Load a server 🤔">
            <a v-on:click="ResetServer(TmpServer)" class="button is-small is-rounded">GO !</a>
            </p>
          </div>
        </div>
      </div>

      <p>
        Current Server <span class="tag is-primary">{{ MyServer }}</span>
      </p>
    </h2>
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
    <section class="hero is-small">
      <div class="hero-body">
        <div class="has-text-centered">
          <a v-on:click="StartMyApi" class="button is-primary">Launch My Api</a>
        </div>
      </div>
    </section>
  </div>

</template>

<script>
  export default {
    data() {
      const code = ` HELLO WORLD LOL`
      return {
        code,
        MyServer : 'none' ,
        TmpServer: '',
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
        console.log('onCmReady', codemirror.setSize(null, 500))
      },
      onCmFocus(codemirror) {
        console.log('onCmFocus', codemirror)
      },
      onCmBlur(codemirror) {
        console.log('onCmBlur', codemirror)
      },
      ResetServer(NewServer){
        this.MyServer = NewServer
        this.StartMyApi()
      },
      async  StartMyApi() {
        let data = JSON.stringify({
          code: this.code,
          server: this.MyServer,
        })

        console.log(data)

        this.$axios.$post('http://localhost:8077/api?'+ data, {
          headers: { 'Content-Type': 'application/json',}}).then(function (response) {
          console.log(response);
        })



      }
    }
  }
</script>

<style lang="scss" scoped>


  .CodeMirror {
    border: 1px solid #eee;
    height: auto;
  }

  .tiny-input{
    width: 140px;
  }

</style>
