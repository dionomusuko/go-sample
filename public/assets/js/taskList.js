const taskList = new Vue({
    el: '#task-list',
    data: {
        user: '',
        newTask: '',
        tasks: [],
    },
    methods: {
        getTasks() {
            const url = 'api/tasks'
            const headers = {'Authorization': `Bearer ${this.getToken()}`}

            fetch(url, {headers}).then(response => {
                if(response.ok) {
                    return response.json()
                }
                return []
            }).then(json => {
                this.tasks = json
            })
        },
        postTask() {
            const url = 'api/tasks'
            const method = 'POST'
            const headers = {
                'Authorization': `Bearer ${this.getToken()}`,
                'Content-Type': 'application/json; charset=UTF-8'
            }
            const body = JSON.stringify({name: this.newTask})

            fetch(url, {method, headers, body}).then(response => {
                if(response.ok) {
                    return response.json()
                }
            }).then(json => {
                if(typeof json === 'undefined') {
                    return
                }
                this.tasks.push(json)
                this.newTask = ''
            })
        },
        deleteTask(id) {
            const url = `api/tasks/${id}`
            const method = 'DELETE'
            const headers = {'Authorization': `Bearer ${this.getToken()}`}

            fetch(url, {method, headers}).then(response => {
                if(response.ok) {
                    this.tasks = this.tasks.filter(task => task.id !== id)
                }
            })
        },
    created() {
        const date = new Date()
        const claims = JSON.parse(atob(this.getToken().split('.')[1]))
        this.user = claims.name
        if(claims.exp < Math.floor(date.getTime() / 1000)) {
            this.logout()
        } else {
            this.getTasks()
        }
    },
})