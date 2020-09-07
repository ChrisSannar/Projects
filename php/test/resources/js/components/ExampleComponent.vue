<template>
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Example Component</div>

                    <div class="card-body">
                        <p v-for="task of tasks" :key="task.id">
                            {{ task.text }} - {{ task.completed ? 'ok' : 'x' }}
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            tasks: []
        }
    },
    mounted() {
        console.log('Component mounted.');
        this.getTasks()
            .then((resp) => {
                let { data } = resp.data;
                this.tasks = data;
            })
            .catch(console.error);
    },
    methods: {
        getTasks: () => {
            return axios.get('/test')
        }
    }
}
</script>
