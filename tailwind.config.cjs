/** @type {import('tailwindcss').Config}*/
const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			fontFamily: {
				obviously: ['obviously'],
				obviouslywide: ['obviously-wide'],
				obviouslycondensed: ['obviously-condensed']
			},
			colors: {
				purple: '#242038',
				lilac: '#C297B8',
				pink: '#FDCFF3',
				darkpink: '#DE89BE'
			},
			boxShadow: {
				'white-depth': '0px 3px 0px 0px rgba(255, 255, 255, 0.69)'
			}
		}
	},

	plugins: []
};

module.exports = config;
