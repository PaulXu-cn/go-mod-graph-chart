const path = require('path');
const webpack = require('webpack');




/*
 * SplitChunksPlugin is enabled by default and replaced
 * deprecated CommonsChunkPlugin. It automatically identifies modules which
 * should be splitted of chunk by heuristics using module duplication count and
 * module category (i. e. node_modules). And splits the chunks…
 *
 * It is safe to remove "splitChunks" from the generated configuration
 * and was added as an educational example.
 *
 * https://webpack.js.org/plugins/split-chunks-plugin/
 *
 */

/*
 * We've enabled TerserPlugin for you! This minifies your app
 * in order to load faster and run less javascript.
 *
 * https://github.com/webpack-contrib/terser-webpack-plugin
 *
 */

const TerserPlugin = require('terser-webpack-plugin');


/*
 * webpack Html Plugin
 */
const HtmlWebpackPlugin = require('html-webpack-plugin');


/*
 * 引入 ParallelUglifyPlugin 插件
 */
const ParallelUglifyPlugin = require('webpack-parallel-uglify-plugin');


/*
 * 引入自定义插件
 */
const PackAllInGoPlugin = require('./plugin/pack-all-in-go-plugin');


const isDev = process.env.NODE_ENV === 'development'


config = {
  mode: 'development',
  devtool: false, // 编译成果 保留换行
  plugins: [
    new HtmlWebpackPlugin({
      title: 'go mod graph charting',
      filename: 'index.html',
      template: './src/temp-index.html'
    }),
    new ParallelUglifyPlugin({
      // 传递给 UglifyJS的参数如下：
      uglifyJS: {
        output: {
          /*
           是否输出可读性较强的代码，即会保留空格和制表符，默认为输出，为了达到更好的压缩效果，
           可以设置为false
          */
          beautify: true,
          /*
           是否保留代码中的注释，默认为保留，为了达到更好的压缩效果，可以设置为false
          */
          comments: false
        },
        compress: {
          /*
           是否在UglifyJS删除没有用到的代码时输出警告信息，默认为输出，可以设置为false关闭这些作用
           不大的警告
          */
          warnings: false,

          /*
           是否删除代码中所有的console语句，默认为不删除，开启后，会删除所有的console语句
          */
          drop_console: true,

          /*
           是否内嵌虽然已经定义了，但是只用到一次的变量，比如将 var x = 1; y = x, 转换成 y = 5, 默认为不
           转换，为了达到更好的压缩效果，可以设置为false
          */
          collapse_vars: true,

          /*
           是否提取出现了多次但是没有定义成变量去引用的静态值，比如将 x = 'xxx'; y = 'xxx'  转换成
           var a = 'xxxx'; x = a; y = a; 默认为不转换，为了达到更好的压缩效果，可以设置为false
          */
          reduce_vars: true
        },
        beautify: {
          semicolons: false
        }
      },
      test: /.js$/g,
      include: [],
      exclude: [],
      cacheDir: '',
      workerCount: '',
      sourceMap: false
    }),
    new webpack.ProgressPlugin(),
    new PackAllInGoPlugin({options: true})
  ],

  module: {
    rules: [{
      test: /.css$/,

      use: [{
        loader: "style-loader"
      }, {
        loader: "css-loader",

        options: {
          sourceMap: true
        }
      }]
    }]
  },

  optimization: {
    minimizer: [new TerserPlugin({
      terserOptions: {
        comments: false
      },
      extractComments: false
    })],

    splitChunks: {
      cacheGroups: {
        vendors: {
          priority: -10,
          test: /[\\/]node_modules[\\/]/
        }
      },

      chunks: 'async',
      minChunks: 1,
      minSize: 30000,
      name: false
    }
  }
}


if (isDev) {
  config.devServer = {
    host: '0.0.0.0',
    contentBase: path.join(__dirname, 'src', 'static'),
    open: true,
    port: 8080,
    overlay: {
      errors: true
    }
  }
} else {
}


module.exports = config
