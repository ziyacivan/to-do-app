# GitHub Copilot Custom Instructions for Todo Application

## Commit Messages
When you are generating commit messages, use conventional commit messages structure (semantic)

## Technical Context
I'm building a Todo application using an event-driven microservices architecture with Kafka. This project demonstrates system design patterns in a distributed environment. The application consists of separate services that communicate through events.

## Tech Stack
- **Language**: Go 1.21+
- **Messaging**: Apache Kafka 3.5+
- **Databases**: PostgreSQL 15+ (separate instances per service)
- **HTTP Framework**: Gin/Gonic
- **ORM**: GORM
- **API Format**: JSON REST
- **Containerization**: Docker 24.0+

## Architecture Components
- Task Service (core service managing todo items)
- Notification Service (consumes events and handles notifications)
- Kafka (event broker connecting services)
- PostgreSQL databases (separate for each service)

## Event Types
- task_created: {id, title, timestamp}
- task_updated: {id, new_status}
- task_deleted: {id, reason}

## Coding Style Preferences
- Use idiomatic Go with standard project layout
- Implement clean separation of concerns (handlers, services, repositories)
- Include appropriate error handling and logging (using Zap)
- Write testable code with high test coverage (target <3% coverage gap)
- Use dependency injection patterns for testability
- Follow SOLID principles

## Performance Requirements
- API Response Time: <500ms (p95)
- Event Processing Latency: <1s (p99)
- Support throughput of 1000 req/sec
- Implement at-least-once event delivery semantics
- Include dead letter queue for failed messages

## Security Considerations
- HTTPS for API endpoints
- Database credential encryption
- JWT authentication (for future implementation)

## When helping me:
1. Suggest code that follows the event-driven architecture pattern
2. Prioritize scalability and reliability in your solutions
3. Include proper error handling for both application and network errors
4. Remember each service has its own database
5. Consider Kafka best practices for event ordering and delivery guarantees
6. Suggest appropriate monitoring and instrumentation code

## Related Documentation References
- Confluent Kafka Documentation
- Go Microservices Patterns
- PostgreSQL ACID Compliance
