<template>
    <div>
        <div>
            <label for="email">
                <span> email</span>
                <input  
                    v-bind:value="email" 
                    @input="email = $event.target.value" 
                    type="text" 
                    name="email" 
                    placeholder="E-mail"
                    >
            </label>
            <br>
            <label for="pass">
                <span> Пароль </span>
                <input 
                    v-bind:value="password" 
                    @input="password = $event.target.value" 
                    type="text" 
                    name="pass" 
                    placeholder="bla baka">
            </label>
        </div>

        <div>
            <label for="headerLetter">
                <span> Заголовок сообщения</span>
                <input 
                    v-bind:value="headerLetter" 
                    @input="headerLetter = $event.target.value" 
                    type="text" 
                    name="headerLetter" 
                    placeholder="header">
            </label>
            <br>
            <label for="bodyLetter">
                <span>Текст сообщения</span>
                <input 
                v-bind:value="bodyLetter" 
                @input="bodyLetter = $event.target.value" 
                type="text" 
                name="bodyLetter"
                placeholder="body"
                >
            </label>

        </div>

        <div>
            <label for="emails">
                <span>Адреса электронных почт через пробел без дополнительных знаков.</span>
                <br>
                <textarea 
                v-bind:value="emails"
                @input="emails = $event.target.value" 
                name="emails" 
                cols="30" 
                rows="10"
                ></textarea>
            </label>
        </div>
        <div>
            <button @click="viewLetterData">Отправить</button>
        </div>
    </div>
</template>
    


<script>
import axios from 'axios';
export default {
    data() {
        return {
            email: "",
            password: "",
            headerLetter: "",
            bodyLetter: "",
            emails: ""
        }
    },
    methods: {
        viewLetterData(){
            const newLetter = {
                letter: {
                    header: this.headerLetter,
                    body: this.bodyLetter,
                },
                sender: {
                    email: this.email,
                    password: this.password
                },
                recipients: this.emails
            } 

            console.log(this.emails)
            axios.post('http://localhost:8080/api/v1/sendLetters', newLetter)
            .then((response) => {
                console.log(response);
            })
            .catch((error) => {
                console.log(error);
            });
            console.log(newLetter)
        }
    }
}

</script>

<style>

</style>