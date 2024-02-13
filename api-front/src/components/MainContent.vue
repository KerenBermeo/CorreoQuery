<script setup lang="ts">
import { defineProps } from 'vue';

// Definir las propiedades del componente
const props = defineProps<{
  email: { from: string; to: string[]; subject: string; date: string; content: string };
}>();
console.log(props.email.to[0])
</script>

<template>
    <!-- Contenido principal -->
    <div class="main-content overflow-y-auto px-4 items-center">
        <div class="email-view  bg-white  rounded-lg p-8">
          <!-- Detalles del correo -->
          <h2 class="text-2xl font-semibold mb-4">{{ props.email.subject }}</h2>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <span class="font-semibold text-gray-700">From:</span>
              <div class="text-gray-800">{{ props.email.from }}</div>
            </div>
            <div  class="col-span-2 overflow-y-auto max-width" >
              <span class="font-semibold text-gray-700 ">To:</span>
              <div class="text-gray-800">                
                <span v-for="(recipient, indexi) in props.email.to" :key="indexi">                                
                    <span v-for="(toEmail, indexj) in recipient" :key="indexj">
                        {{ toEmail }} - 
                    </span>                    
                </span>    
              </div>
            </div>
            <div>
              <span class="font-semibold text-gray-700">Subject:</span>
              <div class="text-gray-800">{{ props.email.subject }}</div>
            </div>
            <div>
              <span class="font-semibold text-gray-700">Date:</span>
              <div class="text-gray-800">{{ props.email.date }}</div>
            </div>
            <div class="col-span-2 overflow-y-auto max-h-48">
            <span class="font-semibold text-gray-700">Message:</span>
                <div class="text-gray-800">
                <span v-for="(recipient, indexi) in props.email.content" :key="indexi">
                    {{ recipient }}
                </span>
                <span v-for="(container, indexA) in props.email.to" :key="indexA">                                
                    <span v-for="(contentEmail, indexB) in container" :key="indexB">
                        {{ contentEmail }}, <br>
                    </span>                    
                </span>
                </div>
            </div>
          </div>
        </div>
    </div>
</template>

<style scoped>
.main-content {
  width: 969px;
  height: 400px;
}
@media (max-width: 767px) {
  .main-content {
    width: 416px;
    height: 330px; /* Cambia a una sola columna */
  }
}


</style>


  