# API Design Interview

## Goal

Assess whether you can design clear API contracts and reason about client/server behavior.

Based on the expected interview format, focus on API design rather than broad system design.

## Likely Prompts

- Design APIs for user login and registration.
- Design APIs for todo/task management.
- Design APIs for file upload and download.
- Design APIs for search with filtering, sorting, and pagination.
- Design APIs for notifications.
- Design APIs for audit logs.
- Design APIs for case or evidence management.
- Design APIs for role-based access control.
- Design APIs for report generation.
- Design APIs for webhook/event subscription.

## Recommended Answer Structure

1. Clarify users, resources, and main use cases.
2. Define endpoints and HTTP methods.
3. Define request and response bodies.
4. Explain validation rules and error format.
5. Choose status codes.
6. Add pagination, filtering, and sorting where needed.
7. Cover authentication, authorization, and privacy.
8. Handle idempotency and retries for create operations.
9. Mention tests and backward compatibility.

## Evidence Upload Example

> I would first clarify who can upload evidence, what metadata is required, file size limits, and whether upload should be direct or pre-signed. A simple contract could be `POST /v1/evidence/uploads` to create an upload session, then `PATCH /v1/evidence/{id}` to update metadata. I would include validation errors in a consistent format, require authorization by case access, and keep audit logs for create/update actions. For retries, I would support an idempotency key so the client does not create duplicate evidence records.

Tie to your CV:

> In my current role, I designed backend APIs for payment and order workflows where correctness, idempotency, state transitions, and auditability mattered. I would bring the same careful API thinking to Axon's mission-critical workflows.
