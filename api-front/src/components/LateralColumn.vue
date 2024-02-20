<script setup lang="ts">
import { ref } from 'vue';
import MainContent from "./MainContent.vue";
import type {Email} from '../stores/email';

// Definir las propiedades del componente
const emails = ref<Email[]>([]);

// Correo electrónico seleccionado
const selectedEmail = ref<Email | null>(null);


// Método para obtener la lista de correos electrónicos
const fetchEmails = async () => {
  try {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/api/emails`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({})
    });
    if (!response.ok) {
      throw new Error('Error al cargar los correos electrónicos: ' + response.statusText);
    }
    const data = await response.json();
   
    // console.log(data.hits.hits);

    emails.value = data.hits.hits.map((email: any) => ({
      id: email._id,
      from: email._source.from,
      date: email._source.date,
      to: email._source.to,
      subject: email._source.subject,
      content: email._source.content
    }));

   

    // Mostrar el primer correo en el contenido principal por defecto
    if (emails.value.length > 0) {
      selectEmail(emails.value[0]);
    }
   

  } catch (error) {
    console.error(error);
  }
};

// Método para seleccionar un correo y mostrarlo en el contenido principal
const selectEmail = (email: Email) => {
  selectedEmail.value = email;
};

// Llamar al método para obtener la lista de correos electrónicos
fetchEmails();
</script>

<template>
  <div class="email-app mt-16 h-screen flex items-center justify-center bg-gray-900 text-white">
    <!-- Barra lateral -->
    <div class="sidebar w-64 bg-gray-800 shadow-xl rounded-lg border-r border-gray-700 h-3/5">
      <!-- Título -->
      <div class="p-4 overflow-x-auto">
        <h1 class="text-2xl font-semibold">From & Date</h1>
      </div>
      <!-- Lista de correos -->
      <div class="border-t border-gray-700 h-4/5 overflow-y-auto scrollbar-thin scrollbar-thumb-gray-500 scrollbar-track-gray-300">
        <ul class="space-y-2">
          <!-- Iterar sobre los remitentes y asuntos de correo y mostrar cada uno -->
          <li v-for="email in emails" :key="email?.id">
              <a href="#" class="email-item block py-2 px-4 hover:bg-gray-700 transition duration-300" @click="selectEmail(email)">
                <span class="font-semibold"><span class="font-bold">{{ email.from }}</span> - <span class="font-thin">{{ email.date }}</span></span>
              </a>
          </li>
        </ul>
      </div>
    </div>
    <!-- Contenido principal -->
    <MainContent :email="selectedEmail"/>

  </div>
</template>



<style scoped>
/* Estos estilos son para navegadores WebKit (Chrome, Safari) */
::-webkit-scrollbar {
  width: 6px; /* Ancho de la barra de desplazamiento */
}

::-webkit-scrollbar-thumb {
  background-color: #4A5568; /* Color del relleno de la barra de desplazamiento */
  border-radius: 3px; /* Radio de borde de la barra de desplazamiento */
}

::-webkit-scrollbar-track {
  background-color: #CBD5E0; /* Color del fondo de la barra de desplazamiento */
}

@media (max-width: 767px) {
  .email-app {
    flex-direction: column; /* Cambia a una sola columna */
  }
}
</style>

  
  
