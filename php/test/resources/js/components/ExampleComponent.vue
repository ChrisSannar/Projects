<template>
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Example Component</div>

                    <div class="card-body">
                        <div v-for="(task, index) of tasks" :key="index">
                            <input 
                                type="checkbox" 
                                name="completed" 
                                :id="'completed-' + index" 
                                :checked="!!task.completed" 
                                @change="markTask($event, index)">
                            <input 
                                type="text" 
                                name="text" 
                                :id="'text-' + index" 
                                :value="task.text" 
                                @keydown="changeTask($event, index)">
                        </div>
                        <button @click="saveList">Save</button>
                        <p v-if="message">{{ message }}</p>
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
            tasks: [],
            message: "",
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
        },
        saveTasks: (tasks) => {
            return axios.put('/test', tasks);
        },
        markTask(event, index) {
            // console.log('mark', event.target.checked, index);
            this.tasks[index].completed = event.target.checked ? 1 : 0;
        },
        changeTask(event, index){
            this.tasks[index].text = event.target.value;
        },
        saveList() {
            console.log('saveTasks', this.tasks);

            let self = this;
            this.saveTasks(this.tasks)
                .then((resp) => {
                    console.log("RESP", resp);
                    // self.message = "Tasks Saved";
                    // setTimeout(() => {
                    //     self.message = "";
                    // }, 3000)
                })
                .catch(console.error);
        }
    }
}
</script>
