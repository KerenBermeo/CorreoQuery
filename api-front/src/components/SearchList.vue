<script setup lang="ts">
   import { useRoute} from 'vue-router'
   import { onMounted, ref } from 'vue'
   import type {SearchResult} from '../stores/searchResult'

  const route = useRoute()
  const params = route.params.query
  const result = ref<SearchResult[]>([]);
  
  // Función para realizar la solicitud de búsqueda
  const search = async () => {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/api/makesearch`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          parameter: params,
          num: 1
        })
      });

      if (!response.ok) {
        throw new Error('Error al realizar la búsqueda');
      }

      const data = await response.json();
      result.value = data.hits.hits;
      console.log('Resultados de la búsqueda:', result.value);

      // Aquí puedes manejar la respuesta y mostrarla en tu página
    } catch (error) {
      console.error('Error:', error);
    }
  };

  // Llamar a la función de búsqueda cuando se monta el componente
  onMounted(() => {
    search();
  });

</script>

<template>
  <div class="container mx-auto px-4 py-8 bg-gray-900 text-white min-h-screen">
    <!-- Botón de retroceso -->
    <router-link to="/" class="text-gray-400 hover:text-gray-300 mb-4 inline-block">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 inline-block mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
      </svg>
      Back
    </router-link>

    <!-- Título de los resultados de búsqueda -->
    <h1 class="text-3xl font-bold mb-4">Search Results for "{{ params }}"</h1>

    <!-- Resultados de la búsqueda -->
    <div v-if="result.length > 0">
      <div v-for="email in result" :key="email._id" class="bg-gray-800 rounded-lg shadow-lg mb-6">
        <button class="block w-full bg-gray-700 hover:bg-gray-600 p-6 rounded-lg transition duration-300">
          <div class="flex justify-between items-center">
            <div class="w-1/4">
              <p class="text-lg font-semibold">Subject:  {{ email._source.subject }}</p>
            </div>
            <div class="w-3/4">
              <p class="text-gray-400">From:  {{ email._source.from }}</p>
              <p class="text-gray-400">To:  {{ email._source.to[0] }}</p>
              <p class="text-gray-400">Date:  {{ email._source.date }}</p>
            </div>
          </div>
        </button>
        <div class="overflow-y-auto h-20">
          <h1>Content:</h1> 
          <p class="text-gray-400">Date:  {{ email._source.content }}</p>
        </div>
      </div>
    </div>
    <div v-else>
      <p class="text-gray-400">No results found.</p>
    </div>
  </div>
</template>

  
