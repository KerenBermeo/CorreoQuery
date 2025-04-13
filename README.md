# **EmailQuery**

**EmailQuery** es una interfaz diseñada para buscar eficientemente información en una base de datos de correos electrónicos. Este proyecto permite explorar y recuperar datos relevantes almacenados en la base de datos, facilitando su gestión y búsqueda.

---

## **Descripción**

Este proyecto utiliza el conjunto de datos **Enron Email Dataset**, que contiene más de 500,000 correos electrónicos de la empresa Enron, y permite realizar búsquedas sobre estos correos a través de una interfaz sencilla. El backend fue desarrollado en **Go** y el frontend en **Vue 3** con **TypeScript**. El motor de búsqueda utilizado es **ZincSearch**, una alternativa de código abierto a **Elasticsearch**.

**Características principales:**
- **Búsqueda eficiente:** Consultas rápidas sobre los correos.
- **Despliegue en la nube:** Implementación usando **AWS**.
- **Diseño frontend básico:** Interfaz de usuario con **Vue 3** y **Tailwind CSS**.

---

## **Cómo ejecutar la aplicación**

### **1. Subir los datos de Enron a AWS S3:**

Primero, debes descargar los datos de Enron disponibles en [CMU Enron Dataset](http://www.cs.cmu.edu/~enron) y cargarlos a tu bucket de AWS S3.  
Asegúrate de tener configurado tu entorno de AWS CLI correctamente. Puedes encontrar más información en la [documentación oficial de AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html).

### **2. Ejecutar localmente:**

Si deseas ejecutar la aplicación localmente, sigue estos pasos:

#### **a. Clona el repositorio:**

```bash
git clone https://github.com/KerenBermeo/CorreoQuery.git
cd CorreoQuery
```

#### **b. Configura tus variables de entorno:**

Crea un archivo `.env` en el directorio raíz y configura las variables necesarias para el backend y frontend:

- **Backend (vars.env):**
  ```env
  ZINC_FIRST_ADMIN_USER=admin
  ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123
  ZINC_SEARCH_URL=""
  ZINC_SERVER_NAME_INDEX=desired_name_for_index
  BACK_ROOT_DIRECTORY=path_to_your_data
  REQUEST_ORIGIN=""
  ```

- **Frontend (.env):**
  ```env
  VITE_API_URL=""
  ```

#### **c. Ejecuta el contenedor con Docker:**

```bash
docker-compose up -d
```

---

## **Despliegue en AWS**

Para desplegar la aplicación en **AWS EC2**, sigue estos pasos:

### **1. Instalar Terraform:**

Primero, instala **Terraform** en tu máquina. Puedes seguir los pasos de instalación en su [página oficial](https://www.terraform.io/downloads.html).

### **2. Configurar la infraestructura en AWS:**

Dentro del proyecto, el archivo `main.ts` está configurado para trabajar con AWS. Ejecuta los siguientes comandos para iniciar, planificar y aplicar la infraestructura:

```bash
terraform init
terraform plan
terraform apply
```

### **3. Subir los datos a S3 y ejecutar la aplicación:**

Una vez que la instancia de AWS esté corriendo, conéctate a ella y ejecuta los siguientes pasos:

```bash
git clone https://github.com/KerenBermeo/CorreoQuery.git
cd CorreoQuery
sudo apt install awscli
aws configure
aws s3 cp s3://bucket_name/compressed_file_name /tmp/compressed_file_name
# Extrae los archivos en una carpeta
unzip /tmp/compressed_file_name -d /desired_directory
# Elimina el archivo comprimido
rm /tmp/compressed_file_name
```

#### **4. Configura las variables de entorno en AWS:**

Asegúrate de configurar las variables de entorno en el servidor para que el backend y el frontend funcionen correctamente.

#### **5. Ejecuta el contenedor en AWS:**

```bash
docker-compose up -d
```

---

## **Tecnologías utilizadas**

- **Backend:** Go (ETL, API REST)
- **Frontend:** Vue 3, TypeScript, Tailwind CSS
- **Motor de búsqueda:** ZincSearch (alternativa a Elasticsearch)
- **Despliegue:** Docker, AWS EC2, Terraform
- **Infraestructura como código:** Terraform

---

## **Próximos pasos**

Este proyecto está en constante desarrollo. Las siguientes mejoras están planificadas:

- Optimización del proceso ETL y de concurrencia.
- Mejoras en la experiencia de usuario (UX).
- Refactorización del código siguiendo las buenas prácticas y principios SOLID.

---

## **Licencia**

Este proyecto está bajo la licencia MIT. Ver el archivo [LICENSE](LICENSE) para más detalles.

---

### 🔍 **¿Te gustaría probar la aplicación?**

¡Puedes hacerlo! Si estás interesado en ver cómo funciona el proyecto o quieres colaborar, no dudes en seguir los pasos de instalación y despliegue. 
