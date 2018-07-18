module.exports = {
    /*
     ** Headers of the page
     */
    head: {
        title: 'myjson',
        meta: [
            {
                charset: 'utf-8'
            },
            {
                name: 'viewport',
                content: 'width=device-width, initial-scale=1'
            },
            {
                hid: 'description',
                name: 'description',
                content: 'Nuxt.js project'
            }
    ],
        link: [
            {
                rel: 'icon',
                type: 'image/x-icon',
                href: '/favicon.ico'
            }
    ]
    },
    css: [
    // lib css
    'codemirror/lib/codemirror.css',
    // merge css
    'codemirror/addon/merge/merge.css',
    // theme css
    'codemirror/theme/base16-dark.css',
    'codemirror/theme/blackboard.css',
      { src: '@/assets/api-simulator.scss', lang: 'sass' },
      { src: '@/assets/animate.css', lang: 'css' },
  ],

  modules: [
    '@nuxtjs/axios',
  ],


  plugins: [
        {
            src: '~/plugins/vue-codemirror.js',
            ssr: false
        },
      { src:'~plugins/buefy', ssr: true},
  ],
    /*
     ** Customize the progress bar color
     */



    loading: {
        color: '#3B8070'
    },
    /*
     ** Build configuration
     */
    build: {
        /*
         ** Run ESLint on save
         */
        extend(config, {
            isDev,
            isClient
        }) {
            if (isDev && isClient) {
                config.module.rules.push({
                    enforce: 'pre',
                    test: /\.(js|vue)$/,
                    loader: 'eslint-loader',
                    exclude: /(node_modules)/
                })
            }
        }
    }
}
