var path = require('path');
var webpack = require('webpack');

module.exports = {
	devtool: 'eval',
	entry: './app/app.js',
	output: {
	 path:'',
	 filename: 'bundle.js',
	 publicPath: ''
 },
	//
	resolve: {
		modulesDirectories: [
			'src',
			'node_modules'
		],
		extensions: [ '', '.json', '.js' ,'.jsx', '.css','.jsx']
	},
	module: {
		loaders: [
			{
				test: /\.jsx?$/,
				loader: 'babel',
				exclude: /node_modules/,
				query: {
					presets: ["es2015", "stage-0", "react"]
				}
			},
			{
				test: /\.css$/,
				loader:'style!css!'
			},
			{
				test: /\.json$/,
				loader: 'json-loader'
			},
			{
				test: /\.md/,
				loaders: [ "html-loader", "markdown-loader" ]
			},

		{ test: /\.woff2(\?v=\d+\.\d+\.\d+)?$/,   loader: "url?limit=10000&minetype=application/font-woff2" },
		{ test: /\.woff(\?v=\d+\.\d+\.\d+)?$/,   loader: "url?limit=10000&minetype=application/font-woff" },
		{ test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,    loader: "url?limit=10000&minetype=application/octet-stream" },
		{ test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,    loader: "file" },
		{ test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,    loader: "url?limit=10000&minetype=image/svg+xml" },
		]
	},

	plugins: [
		new webpack.ProvidePlugin({
		    $: "jquery",
		    jQuery: "jquery"
		  }),
		],

}


