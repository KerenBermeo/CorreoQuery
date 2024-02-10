<script setup lang="ts">
  const email = {
      from: 'remitente@example.com',
      to: ['destinatario1@example.com', 'destinatario2@example.com'],
      subject: 'Asunto del correo',
      date: '2024-02-09',
      content: 'Contenido del correo electrónico'
  }
</script>

<template>
    <div class="email-app mt-16 flex items-center justify-between bg-gray-100 h-screen">
      <!-- Barra lateral -->
      <div class="sidebar w-64 bg-white shadow-xl rounded-lg overflow-x-auto custom-scrollbar">
        <div class="p-4">
          <!-- Logo o título -->
          <h1 class="text-2xl font-semibold text-gray-800">Message ID</h1>
        </div>
        <div class="border-t border-gray-300">
          <!-- Lista de correos -->
          <ul class="space-y-2">
            <li>
              <a href="#" class="email-item block py-2 px-4 text-gray-800 hover:bg-gray-200 transition duration-300">
                <span class="font-semibold">Correo 1</span>
              </a>
            </li>
            <!-- Otros correos -->
          </ul>
        </div>
      </div>
      <!-- Contenido principal -->
      <div class="main-content flex-1 px-4">
        <div class="email-view bg-white shadow-xl rounded-lg p-8">
          <!-- Detalles del correo -->
          <h2 class="text-2xl font-semibold mb-4">{{ email.subject }}</h2>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <span class="font-semibold text-gray-700">From:</span>
              <div class="text-gray-800">{{ email.from }}</div>
            </div>
            <div>
              <span class="font-semibold text-gray-700">To:</span>
              <div class="text-gray-800">
                <span v-for="(recipient, index) in email.to" :key="index">{{ recipient }}</span>
              </div>
            </div>
            <div>
              <span class="font-semibold text-gray-700">Subject:</span>
              <div class="text-gray-800">{{ email.subject }}</div>
            </div>
            <div>
              <span class="font-semibold text-gray-700">Date:</span>
              <div class="text-gray-800">{{ email.date }}</div>
            </div>
            <div class="col-span-2">
              <span class="font-semibold text-gray-700">Message:</span>
              <div class="text-gray-800">{{ email.content }}</div>
            </div>
            <!-- Otros detalles del correo -->
          </div>
          <!-- Botones o acciones -->
          <div class="mt-8 flex justify-end">
            <!-- Botones de acción -->
          </div>
        </div>
      </div>
    </div>
  </template>
  
  
  <style scoped>
  /* Estilos específicos para este componente */
  .email-app {
    font-family: 'Roboto', sans-serif;
  }
  
  .sidebar {
    min-width: 280px;
  }
  
  .email-item {
    transition: background-color 0.3s ease;
  }
  
  .email-item:hover {
    background-color: #f0f0f0;
  }
  
  .email-view {
    min-width: 320px;
  }
  
  </style>
  
