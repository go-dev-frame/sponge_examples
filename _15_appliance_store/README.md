## Project Practical Example —— Building a Home Appliance Retail Management Platform from Scratch

Below is an example of building a product management platform for an offline home appliance store, illustrating how to use Sponge and DeepSeek to collaboratively develop backend services. The backend technology stack chosen for this example is **Web Service (Gin + Gorm + Protobuf)**.

> **Tip:** Here, the request and response data structures of the API interfaces are defined in Protobuf files, taking full advantage of Protobuf's benefits—parsing Protobuf to generate the necessary framework code and API interface documentation.

<br>

### 1. Generate Functional Requirements Document

First, generate a detailed functional requirements document using DeepSeek R1. Input the following prompt:

> "Now, we need to implement the backend service for a product management platform of an offline home appliance store. Please list the detailed functional requirements."

DeepSeek R1 will generate a comprehensive requirements document. Developers can remove unnecessary functions based on actual needs, retain the required functional modules, or add additional functional modules. Click to view the [Home Appliance Retail Management Platform Functional Requirements Document](https://github.com/go-dev-frame/sponge_examples/blob/main/_15_appliance_store/docs/requirements-document.md).

<br>

### 2. Generate MySQL Table Structure DDL

Next, generate the DDL for all MySQL table structures based on the functional requirements document. Input the following prompt:

> "Based on the functional requirements document, generate the DDL for all MySQL table structures required by the backend service. The generated SQL should be directly importable into MySQL to create tables, and each column in the tables should include English comments."

DeepSeek R1 will generate the corresponding MySQL table structure DDL based on the requirements document. Developers need to verify whether it fully meets the requirements and make manual adjustments if necessary. Click to view the [Home Appliance Retail Management Platform Table Structure DDL](https://github.com/go-dev-frame/sponge_examples/blob/main/_15_appliance_store/docs/store.sql).

After importing the MySQL table structure DDL into MySQL, it provides data structure support for subsequent code generation. Sponge can generate various module codes based on these table structures, such as CRUD API code and CRUD Protobuf definitions.

<br>

### 3. API Interface Definition

#### 3.1 Generate CRUD API Protobuf

On Sponge's code generation page, select: 【Public】→【Generate Protobuf CRUD Code】, fill in the parameters, and click the 【Download Code】 button to generate the code, as shown below:

![Protobuf CRUD Code](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_store-protobuf.png)

> **Tip:** If there are many generated proto files, it is recommended to merge them into one file because DeepSeek R1 has a limit on the number of files that can be uploaded.

#### 3.2 Generate Custom API Protobuf

Standard CRUD APIs cannot cover all business requirements, so it is necessary to generate custom API Protobuf definitions based on the CRUD API Protobuf file and the functional requirements document.  
Upload the CRUD API proto file and the home appliance retail management platform functional requirements document to DeepSeek R1, and input the following prompt:

```
All MySQL table standard CRUD API Protobuf definitions have been determined. These APIs only cover part of the backend service functions of the home appliance retail management platform. To cover all functions, please supplement custom API Protobuf definitions based on the CRUD API Protobuf definitions and the home appliance retail management platform functional requirements document. The requirements are as follows:
1. Each rpc method must include option (google.api.http).
2. The rpc method and its message fields must include English comments, and the rpc method must describe the logical implementation process in detail (as the basis for AI-generated business logic code).
3. The supplemented APIs need to identify their belonging Protobuf service.
```

> **Tip:** If the rpc method comments in the generated Protobuf description are not detailed enough (e.g., logical implementation process, specified technology stack), they can be manually supplemented and improved.

The generated custom API Protobuf, together with the CRUD API Protobuf, constitutes the complete API interface definition of the service, providing the basis for subsequent Sponge code generation.

<br>

### 4. Create Service Code

Take the Web service as an example (technology stack: Gin + Gorm + Protobuf) for subsequent code generation and integration.

#### 4.1 Generate Service Base Code

On Sponge's code generation page, select: 【Protobuf】→【Create Web Service】, fill in the parameters, and click the 【Download Code】 button to generate the code. As shown below:

![Create Service Based on Protobuf](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_store-http-pb.png)

The generated code package includes the basic framework of the service. After decompressing, enter the code directory.

#### 4.2 Generate CRUD API Code

Similarly, on Sponge's code generation page, select: 【Public】→【Generate Handler CRUD Code】, fill in the parameters, and click the 【Download Code】 button to generate the code, as shown below:

![Generate Handler-PB Code](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_store-handler-pb.png)

Decompress the file and move the generated `api` and `internal` directories to the service code directory.

Open the project code using VS Code or Goland, and manually merge the custom API Protobuf files generated earlier by DeepSeek R1 into the corresponding proto files in the `api/store/v1` directory.  
Then, execute the following command in the project root directory to generate the code:

```bash
make proto
```

> **Tip:** Whenever the proto file is modified, the `make proto` command needs to be re-executed. Code generation can be specified by the proto file name.

<br>

### 5. Business Logic Code Completion

Sponge integrates the DeepSeek API, which can automatically locate the method functions that need to be supplemented with business logic, allowing the AI assistant to generate business logic implementation code. Execute the following command:

```bash
sponge assistant generate --type=deepseek --model=deepseek-reasoner --api-key=xxxx --dir=.
```

The generated business logic code will be saved in the corresponding directory with a `.assistant` suffix. Developers only need to copy it into the corresponding method functions. The reason for not using automatic filling here is that DeepSeek R1 output may include markdown and other non-pure Go code, so manual verification and copying are required.

<br>

### 6. Test and Verify API Functions

At this point, most of the code has been automatically generated through the collaboration of Sponge and DeepSeek. If the AI assistant cannot meet the requirements based on detailed prompts to generate business logic implementation code, manual coding is required. The complete web service code egg model is as follows:

![web-http-pb-anatomy](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_web-http-pb-anatomy.png)

Next, developers debug and verify the API functions, start the service:

```bash
make run
```

Use a browser to access the Swagger interface for API debugging: [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html)

![Swagger Debug Interface](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/store-swagger.png)

> This is the [backend service example code](https://github.com/go-dev-frame/sponge_examples/tree/main/_15_appliance_store) generated by Sponge and DeepSeek collaboration.
