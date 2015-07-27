App.Models.Task = Backbone.Model.extend({
	defaults: {
		id: 0,
		title: '',
		body: '',
		done: false,
		created_at: ''
	},

	url: function() {
		url = 'http://unix:/tmp/emru.sock:/lists/';

		return url + this.get('list') + '/tasks/' + this.id;
	},

	toJson: function() {
		return _.omit(this.attributes, 'list');
	}
});

App.Models.List = Backbone.Model.extend({
	initialize: function() {
		this.tasks = new App.Collections.Tasks();
	},

	url: function() {
		url = 'http://unix:/tmp/emru.sock:/lists/';

		return url + this.get('name');
	},

	watch: function() {
		this.fetch({error: this.fErr});
	},

	parse: function(response) {
		tasks = response.tasks;
		for (var i = 0; i < tasks.length; i++) {
			task = new App.Models.Task(tasks[i]);
			task.set('list', this.get('name'));
			this.tasks.add(task);
		}
	},

	fErr: function(model, response) {
		console.log('model fetch err:');
		console.log(response);
	},
});
