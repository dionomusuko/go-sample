const loginForm = new Vue({
    el: '#login-form',
    data: {
        name: '',
        email: '',
        password: '',
    },
    methods: {
        login() {
            const url = '/sign-in'
            const method = 'POST'
            const headers = {
                'Content-Type': 'application/json; charset=UTF-8'
            }
            const body = JSON.stringify({
                name: this.name,
                email: this.email,
                password: this.password,
            })

            fetch(url, {method, headers, body}).then(response => {
                if(response.ok) {
                    return response.json()
                } else {
                    alert('Faild to login. Please retry')
                    this.name = ''
                    this.email = ''
                    this.password = ''
                    return {token: ''}
                }
            }).then(json => {
                const token = json.token
                if(token.length > 0) {
                    localStorage.setItem('token', token)
                    location.href = '/tasks'
                }
            })
        },
    },
})