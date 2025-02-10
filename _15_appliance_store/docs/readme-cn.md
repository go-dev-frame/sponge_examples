## 项目实战示例 —— 从零开始构建家电零售管理平台

下面以构建一个线下家电实体店的产品管理平台为例，说明如何利用 Sponge 与 DeepSeek 协同开发后端服务。本示例后端技术栈选择 **Web 服务 (Gin + Gorm + Protobuf)**。

> **提示：** 这里把 API 接口的请求和返回数据结构定义在 Protobuf 文件中，充分利用 Protobuf 的优势——解析Protobuf来生成框架所需的代码和API接口文档。

<br>

### 1. 生成功能需求文档

首先，通过 DeepSeek R1 生成详细的功能需求文档。输入以下提示：

> “现在需要实现线下家电实体店铺的产品管理平台的后台服务，请列出详细的功能需求。”

DeepSeek R1 会生成一个较为全面的需求文档，开发者可以根据实际需要删减不必要的功能，保留真正需要的功能模块，或额外添加补充功能模块。点击查看[家电零售管理平台功能需求文档](https://github.com/go-dev-frame/sponge_examples/blob/main/_15_appliance_store/docs/requirements-document.md)。

<br>

### 2. 生成 MySQL 表结构 DDL

接下来，根据功能需求文档生成所有 MySQL 表结构的 DDL。输入以下提示：

> “根据功能需求文档，生成后台服务所需的所有 MySQL 表结构的 DDL，要求生成的 SQL 可直接导入 MySQL 创建表，表的每列均需附带中文注释。”

DeepSeek R1 会根据需求文档生成对应的 Mysql 表结构 DDL，开发者需要校验判断是否完全满足要求，如果不满足可以人工调整。点击查看[家电零售管理平台表结构DDL](https://github.com/go-dev-frame/sponge_examples/blob/main/_15_appliance_store/docs/store.sql)。

把 Mysql 表结构 DDL 导入 MySQL 后，即可为后续代码生成提供数据结构支持。Sponge 可以根据这些表结构生成各种模块代码，如 CRUD API 代码 和 CRUD Protobuf 定义等。

<br>

### 3. API 接口定义

#### 3.1 生成 CRUD API Protobuf

在 Sponge 的生成代码页面中，依次选择：【Public】→【生成 Protobuf CRUD 代码】，填写参数后点击【下载代码】按钮生成代码，如下图所示：

![Protobuf CRUD 代码](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/store-protobuf.png)

> **提示：** 如果生成的 proto 文件较多，建议将它们合并到一个文件中，因为 DeepSeek R1 上传文件数量有限制。

#### 3.2 生成自定义 API Protobuf

标准的 CRUD API 并不能涵盖所有业务需求，因此需要根据 CRUD API Protobuf 文件和功能需求文档生成自定义 API 的 Protobuf 定义。  
在 DeepSeek R1 中上传 CRUD API 的 proto 文件和家电零售管理平台功能需求文档，并输入如下提示：

```
已确定所有 MySQL 表的标准 CRUD API 的 Protobuf 定义，这些 API 仅涵盖家电零售管理平台后台服务的一部分功能，为了涵盖所有的功能，请依据CRUD API 的 Protobuf 定义和家电零售管理平台的功能需求文档，进行补充自定义API Protobuf，要求如下：
1. 每个 rpc 方法必须包含 option (google.api.http)。
2. rpc 方法及其 message 字段需附带中文注释，rpc 方法需详细描述逻辑实现过程（作为 AI 生成业务逻辑代码的依据）。
3. 补充的 API 需要标识其所属的 Protobuf service。
```

> **提示：** 若生成结果中 Protobuf 描述里的 rpc 方法注释不够详细(例如逻辑实现过程、指定技术栈)，可以适当人工补充完善。

生成的自定义 API Protobuf 与 CRUD API Protobuf 共同构成了服务完整功能的API接口定义，为后续 Sponge 提供生成代码依据。

<br>

### 4. 创建服务代码

以 Web 服务为例（技术栈：Gin + Gorm + Protobuf）进行后续代码生成和集成。

#### 4.1 生成服务基础代码

在 Sponge 代码生成的页面中，选择：【Protobuf】→【创建 Web 服务】，填写参数后点击【下载代码】按钮生成代码。如下图所示：

![基于protobuf创建服务](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/store-http-pb.png)

生成的代码包中包含服务的基本框架，解压后进入代码目录。

#### 4.2 生成 CRUD API 代码

同样在 Sponge 代码生成的页面中，选择：【Public】→【生成 Handler CRUD 代码】，填写参数后点击【下载代码】按钮生成代码，如下图所示：

![生成handler-pb代码](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/store-handler-pb.png)

解压文件，将生成的 `api` 和 `internal` 目录移动至服务代码目录中。

使用 VS Code 或 Goland 打开项目代码，并将在前面由 DeepSeek R1 生成的自定义 API 的 Protobuf 文件人工合并到 `api/store/v1` 目录下对应的 proto 文件中。  
接着，在项目根目录下执行以下命令生成代码：

```bash
make proto
```

> **提示：** 每当修改 proto 文件后，都需重新执行 `make proto` 命令，可通过指定 proto 文件名生成代码。

<br>

### 5. 业务逻辑代码补全

Sponge 集成了 DeepSeek API，可自动定位到需要补充业务逻辑的方法函数，让 AI 助手生成业务逻辑实现代码，执行如下命令：

```bash
sponge assistant generate --type=deepseek --model=deepseek-reasoner --api-key=xxxx --dir=.
```

生成的业务逻辑代码会以 `.assistant` 为后缀保存在相应目录下，开发者只需将其复制到对应方法函数中即可。这里不采用自动填充的方式，是因为 DeepSeek R1 输出可能包含 markdown 等非纯 Go 代码，故需人工校验复制。

<br>

### 6. 测试和验证 API 功能

至此，通过 Sponge 与 DeepSeek 的协同工作，绝大部分代码均已自动生成。如果AI助手根据详细的提示生成业务逻辑实现代码也无法满足要求，则需要人工编写代码。完整的web服务代码鸡蛋模型如下：

![web-http-pb-anatomy](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/web-http-pb-anatomy.png)

接着开发者调试与验证 API 功能，启动服务：

```bash
make run
```

使用浏览器访问 Swagger 界面进行 API 调试：http://localhost:8080/apis/swagger/index.html

![Swagger 调试界面](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/store-swagger.png)

> 这是 Sponge 与  DeepSeek协同生成的[后端服务示例代码](https://github.com/go-dev-frame/sponge_examples/tree/main/_15_appliance_store)。
 
