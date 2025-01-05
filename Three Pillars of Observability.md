# Three Pillars of Observability

## Tracing

### **What is Tracing (Distributed Tracing)?**
   - **Definition**:
     - Tracing is a technique used to **track and visualize the flow of a request** as it moves through a distributed system.
     - It captures the **path of a request** across multiple services, components, and boundaries (e.g., microservices, databases, APIs).

   - **Purpose**:
     - Tracing helps developers and operators understand how a request is processed, identify bottlenecks, and debug issues in complex systems.



### **Key Components of Tracing**
   - **Trace**:
     - A **trace** is the complete record of a request’s journey through the system. It includes all the spans associated with the request.

   - **Span**:
     - A **span** represents a single operation within a trace. It contains:
       - **Start and end timestamps**: To measure the duration of the operation.
       - **Metadata**: Such as the service name, operation name, and tags (key-value pairs for additional context).

   - **Trace ID**:
     - A unique identifier that ties all spans related to a single request together.

   - **Parent-Span ID**:
     - Each span (except the root span) has a **parent-span ID** that links it to the previous operation in the trace.


### **How Tracing Works**
   - **Request Propagation**:
     - When a request enters the system, a unique **trace ID** is generated.
     - As the request propagates through different services, each service adds **spans** to the trace. A span represents a unit of work (e.g., a function call, database query, or API call).

   - **Directed Acyclic Graph (DAG)**:
     - The trace is represented as a **DAG**, where:
       - **Nodes**: Represent spans (units of work).
       - **Edges**: Represent the flow of the request between spans.
     - The DAG structure ensures that the trace is **acyclic** (no loops) and **directed** (shows the sequence of operations).


### **Tools for Distributed Tracing**
   - **OpenTelemetry**: A vendor-neutral framework for collecting traces, metrics, and logs.
   - **Jaeger**: An open-source distributed tracing system.
   - **Zipkin**: A distributed tracing system for collecting and visualizing traces.
   - **AWS X-Ray**: A tracing service for applications running on AWS.
   - **Google Cloud Trace**: A tracing service for applications running on Google Cloud.

### **Example of a Trace**
Imagine a request to an e-commerce application that involves the following steps:
1. **Frontend Service**: Receives the request and forwards it to the **Product Service**.
2. **Product Service**: Queries the **Inventory Service** to check product availability.
3. **Inventory Service**: Queries the **Database** to retrieve inventory data.

The trace for this request might look like this:

```plaintext
Trace ID: abc123
Spans:
- Span 1: Frontend Service (Start: 10:00:00, End: 10:00:02)
  - Span 2: Product Service (Start: 10:00:01, End: 10:00:03)
    - Span 3: Inventory Service (Start: 10:00:02, End: 10:00:04)
      - Span 4: Database Query (Start: 10:00:03, End: 10:00:04)
```

This trace can be visualized as a DAG:
```
Frontend Service (Span 1)
       |
Product Service (Span 2)
       |
Inventory Service (Span 3)
       |
Database Query (Span 4)
```



## Metrics

## Logging
