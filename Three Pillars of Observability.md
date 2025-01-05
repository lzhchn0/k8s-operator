# Three Pillars of Observability

## Tracing

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


### **Tools for Distributed Tracing**
   - **OpenTelemetry**: A vendor-neutral framework for collecting traces, metrics, and logs.
   - **Jaeger**: An open-source distributed tracing system.
   - **Zipkin**: A distributed tracing system for collecting and visualizing traces.
   - **AWS X-Ray**: A tracing service for applications running on AWS.
   - **Google Cloud Trace**: A tracing service for applications running on Google Cloud.

## Metrics

## Logging
