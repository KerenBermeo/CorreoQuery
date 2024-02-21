<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter();

function handleSubmit() {
    const searchElement = document.getElementById('search-navbar') as HTMLInputElement;
    if (searchElement !== null) {
        const searchValue = searchElement.value;
        // Asegúrate de que router no sea undefined antes de usarlo
        if (router) {
            router.push({ name: 'searchList', params: { query: searchValue } });
        } else {
            console.error("Router no está definido");
        }
    } else {
        console.error("Elemento de búsqueda no encontrado en el DOM");
    } 
}
</script>



<!-- Componente de barra de navegacion sacado de https://flowbite.com/docs/components/navbar/ -->
<template>
        <!-- Barra de navegación -->
    <nav class="bg-white border-gray-200 dark:bg-gray-900 fixed w-full top-0">


    <!-- Contenedor principal -->
    <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">

    <!-- Contenido del lado derecho -->
    <div class="flex md:order-2">

        <!-- Botón de búsqueda (visible en dispositivos móviles) -->
        <button id="searchButton" type="button" data-collapse-toggle="navbar-search" aria-controls="navbar-search" aria-expanded="false" class="md:hidden text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 rounded-lg text-sm p-2.5 me-1">
        <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
        </svg>
        <span class="sr-only">Search</span>
        </button>

        <!-- Campo de búsqueda (visible en versiones de escritorio) -->
        <div class="relative hidden md:block">
            <form @submit.prevent="handleSubmit">
                <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                    <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
                    </svg>
                    <span class="sr-only">Search</span>
                </div>
                <input type="text" id="search-navbar" class="block w-full p-2 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search...">
                <button type="submit" class="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-500 hover:text-gray-700 focus:outline-none">
                    <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 22l-6-6-1.999 2L16 18l6-6-6-6-2 2 6 6z"/>
                    </svg>
                </button>
            </form>
        </div>
    </div>

    <!-- Contenido del lado izquierdo (visible en versiones de escritorio) -->
    <div class="items-center justify-between hidden w-full md:flex md:w-auto md:order-1" id="navbar-search">

        <!-- Campo de búsqueda (visible en versiones de escritorio) -->
        <div class="relative mt-3 md:hidden">
            <form @submit.prevent="handleSubmit">
                <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                    <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
                    </svg>
                </div>
                <input type="text" id="search-navbar" class="block w-full p-2 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search...">
                <button type="submit" class="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-500 hover:text-gray-700 focus:outline-none">
                    <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 22l-6-6-1.999 2L16 18l6-6-6-6-2 2 6 6z"/>
                    </svg>
                </button>
            </form>
        </div>

        <!-- Lista de elementos de navegación -->
        <ul class="flex flex-col p-4 md:p-0 mt-4 font-medium border border-gray-100 rounded-lg bg-gray-50 md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
        <li>
            <a href="#" class="block py-2 px-3 text-white bg-blue-700 rounded md:bg-transparent md:text-blue-700 md:p-0 md:dark:text-blue-500" aria-current="page">Home</a>
        </li>
        </ul>
    </div>
    </div>
    </nav>
</template>


