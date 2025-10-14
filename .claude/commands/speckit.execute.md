---
description: Execute specific tasks with automated testing and session documentation
---

## User Input

```text
$ARGUMENTS
```

You **MUST** parse the user input to extract the task range (e.g., "T117-T123", "117-123", or "T117").

## Outline

### 1. Parse Task File and Task Range

Extract task file and task IDs from user input:

**Task File Parsing:**
- Accept formats: "@tasks_02.md T148", "tasks_02.md T148-T150", or just "T148"
- If input starts with "@" or matches pattern "tasks*.md", extract task filename
- Remove "@" prefix if present
- Default to "tasks.md" if no task file specified
- Store as: `TASK_FILE` (e.g., "tasks.md", "tasks_02.md")

**Task Range Parsing:**
- Accept formats: "T117-T123", "117-123", "T117", "117"
- Parse start task ID (remove 'T' prefix if present)
- Parse end task ID if range provided, otherwise use start task as single task
- Store as: `START_TASK` and `END_TASK` (numeric values)

**Examples:**
- "T148" → TASK_FILE="tasks.md", START_TASK=148, END_TASK=148
- "@tasks_02.md T148" → TASK_FILE="tasks_02.md", START_TASK=148, END_TASK=148
- "tasks_02.md T117-T123" → TASK_FILE="tasks_02.md", START_TASK=117, END_TASK=123
- "T105" → TASK_FILE="tasks.md", START_TASK=105, END_TASK=105

### 1b. Documentation Policy (Enforced Throughout Execution)

**CRITICAL**: Only ONE documentation file may be created during execution:

**Permitted Documentation:**
- Path: `docs/{SPEC_ID}/session-summary/t{START_TASK}-t{END_TASK}-summary.yaml`
- Format: `t###-t###-summary.yaml` (lowercase 't', hyphen-separated, numeric range only)
- Example: `t117-t123.yaml` ✓ | `t136-t136.yaml` ✓ (single task)

**Prohibited Documentation (ALL will be deleted):**
- ❌ Descriptive filenames: `t136-transparency-modal-summary.yaml`
- ❌ Files outside `session-summary/` directory
- ❌ Test reports, implementation guides, checklists in `docs/{SPEC_ID}/`
- ❌ `README.md` or any other documentation files
- ❌ Files with feature names in the filename

**Enforcement:**
- This policy applies to ALL agents: orchestrator, test-writer-fixer, code-documenter, coding agents
- Violations detected at the end of execution are automatically deleted
- Only the single permitted session summary survives

### 2. Load Feature Context

Run `.specify/scripts/bash/check-prerequisites.sh --json --require-tasks --include-tasks` from repo root:
- Parse `FEATURE_DIR` (absolute path to current spec folder)
- Parse `AVAILABLE_DOCS` list
- Extract spec ID from FEATURE_DIR (e.g., "/path/to/specs/002-auth" → "002-auth")
- Store as: `SPEC_ID` (e.g., "002-auth")

### 3. Check Checklists Status (Quality Gate)

**If FEATURE_DIR/checklists/ exists:**

1. **Scan all checklist files** in the checklists/ directory
2. **For each checklist**, count:
   - Total items: All lines matching `- [ ]` or `- [X]` or `- [x]`
   - Completed items: Lines matching `- [X]` or `- [x]`
   - Incomplete items: Lines matching `- [ ]`

3. **Create a status table**:
   ```
   | Checklist | Total | Completed | Incomplete | Status |
   |-----------|-------|-----------|------------|--------|
   | ux.md     | 12    | 12        | 0          | ✓ PASS |
   | test.md   | 8     | 5         | 3          | ✗ FAIL |
   | security.md | 6   | 6         | 0          | ✓ PASS |
   ```

4. **Calculate overall status**:
   - **PASS**: All checklists have 0 incomplete items
   - **FAIL**: One or more checklists have incomplete items

5. **If any checklist is incomplete**:
   - Display the table with incomplete item counts
   - **STOP** and ask: "Some checklists are incomplete. Do you want to proceed with implementation anyway? (yes/no)"
   - Wait for user response before continuing
   - If user says "no" or "wait" or "stop", halt execution
   - If user says "yes" or "proceed" or "continue", log warning and proceed to step 4

6. **If all checklists are complete**:
   - Display the table showing all checklists passed
   - Automatically proceed to step 4

**If checklists/ directory does not exist:**
- Skip validation
- Proceed directly to step 4

### 4. Load Task Details and Implementation Context

Read and analyze from FEATURE_DIR:
- **REQUIRED**: `{TASK_FILE}` - Find tasks T{START_TASK} through T{END_TASK}, parse phase structure
- **REQUIRED**: `plan.md` - Tech stack, architecture, file structure
- **OPTIONAL**: `spec.md` - User stories, success criteria (for validation)
- **OPTIONAL**: `data-model.md` - Entity definitions and relationships
- **OPTIONAL**: `contracts/` - API specifications and test requirements
- **OPTIONAL**: `research.md` - Technical decisions and constraints
- **OPTIONAL**: `quickstart.md` - Integration scenarios

**Parse task phases from {TASK_FILE}:**
- Identify phase boundaries (Phase 1: Setup, Phase 2: Foundation, Phase 3+: User Stories, Final Phase: Polish)
- Note which tasks belong to which phase
- Understand phase dependencies and completion criteria
- Store phase information for progress tracking

**For each task in range:**
- Extract task description and phase assignment
- Identify file paths to be modified
- Note parallel execution markers [P]
- Understand dependencies and context
- Determine appropriate agent type (frontend-developer, backend-architect, ai-engineer, etc.)

### 5. Agent Selection Guidelines (Review Before Task Execution)

#### Quick Agent Selection Decision Tree

Apply this decision tree for EVERY task:

```
1. ML/AI features? → ai-engineer
2. New multi-component feature/scaffold? → rapid-prototyper
3. Server/API/database? → backend-architect
4. UI/components/pages? → frontend-developer
5. CI/CD/infrastructure? → devops-automator
6. Mobile native features? → mobile-app-builder
7. Design system? → ui-designer
8. Privacy/compliance? → legal-compliance-checker
9. UX research? → ux-researcher
10. Default: frontend-developer
```

#### Primary Implementation Agents (Choose ONE per task)

**frontend-developer**: React components, Next.js pages, styling, client state
**backend-architect**: API endpoints, database, migrations, auth, server logic
**ai-engineer**: ML models, AI integrations, forecasting, OpenAI API
**rapid-prototyper**: New feature scaffolding, MVPs, proof-of-concepts
**devops-automator**: CI/CD, deployment, monitoring, build optimization
**mobile-app-builder**: React Native, mobile-specific features
**ui-designer**: Design systems, component libraries, visual design
**legal-compliance-checker**: GDPR, privacy controls, audit logging

*See Section 12 for detailed agent descriptions, task indicators, file patterns, and conflict resolution*

---

### 6. Pre-Execution Validation

Before executing tasks, perform validation checks:

**TDD Check:**
1. Identify test tasks in range (descriptions containing "test", "spec", "Test", "Spec", "testing", "unit", "integration")
2. Identify implementation tasks that correspond to those tests
3. If implementation tasks are scheduled before their test tasks:
   - **WARN**: "⚠️ TDD Violation: Implementation task T{n} scheduled before test task T{m}"
   - List affected task pairs
   - Ask user: "Do you want to reorder tasks to follow TDD approach? (yes/no)"
   - If user confirms reordering, adjust execution order
   - If user declines, log warning and proceed as scheduled

**Phase Boundary Check:**
1. Identify which phases the task range spans
2. If tasks span multiple phases:
   - **INFO**: "📦 Task range spans multiple phases: {phase list}"
   - List phase transitions
   - Confirm user wants to execute across phase boundaries
3. If tasks cross from one user story phase to another:
   - **WARN**: "⚠️ Crossing user story boundaries. Ensure previous story is complete."
   - Ask user: "Previous user story phase should be complete. Continue? (yes/no)"

**Dependency Validation:**
1. Check for prerequisite tasks outside the current range
2. If dependencies exist before START_TASK:
   - **WARN**: "⚠️ Tasks have dependencies on earlier tasks not in range"
   - List dependency requirements
   - Ask user: "Prerequisites may not be met. Continue anyway? (yes/no)"

### 7. Execute Tasks Sequentially (Phase-Aware)

**🚨 CRITICAL: ORCHESTRATION-ONLY MODE 🚨**

**YOU ARE AN ORCHESTRATOR. YOU DO NOT IMPLEMENT CODE DIRECTLY.**

**Your ONLY responsibilities:**
1. Parse task requirements
2. Select appropriate agent using decision tree (see Agent Selection Guidelines below)
3. Invoke Task tool with selected agent
4. Track progress and phase transitions

**NEVER:**
- ❌ Write implementation code yourself
- ❌ Edit source files directly
- ❌ Skip agent delegation
- ❌ Combine multiple tasks into one agent call (unless explicitly marked [P] for parallel)
- ❌ Write tests yourself or ask implementation agents to write tests (ONLY test-writer-fixer writes tests)

---

**Track current phase:**
- Monitor which phase each task belongs to
- When entering a new phase, report: "📦 Entering Phase {N}: {Phase Name}"
- When completing a phase, report: "✅ Phase {N} Complete: {task count} tasks finished"
- Validate phase deliverables before advancing to next phase

**For each task T{n} from START_TASK to END_TASK:**

1. **Select agent** (REQUIRED - see Agent Selection Guidelines below):
   - Apply decision tree to choose ONE primary agent
   - Consider task indicators, file patterns, and technical domain
   - If unclear, default to: frontend-developer (UI), backend-architect (API/DB), or rapid-prototyper (new feature)

2. **REQUIRED: Delegate to agent via Task tool**:
   ```
   Task tool with subagent_type: "{selected-agent}"
   Prompt: "
   Task: T{n} - {task description from tasks.md}

   Context:
   - Spec ID: {SPEC_ID}
   - Phase: {current phase name}
   - Related files: {file paths from task analysis}

   Requirements from plan.md:
   {relevant architecture/patterns}

   Success criteria:
   - {criteria from task description}

   CRITICAL CONSTRAINTS:
   - DO NOT create documentation files. Focus only on implementation.
   - DO NOT write tests. Test-writer-fixer agent handles ALL testing in Section 8.
   "
   ```
   - **MANDATORY**: Wait for agent completion before proceeding to next task
   - **MANDATORY**: Do NOT implement the task yourself

3. **Git commit after major changes**:
   - If task represents significant code change (new feature, major refactor, etc.)
   - Create commit with format: `feat({spec-id}): {task description} (T{task-id})`
   - Example: `feat(002-auth): implement password reset endpoint (T117)`
   - Include multiple tasks in one commit if they're tightly related
   - Use bash tool for: `git add . && git commit -m "..."`

4. **Track progress**:
   - Mark task as completed in {TASK_FILE} by changing `- [ ]` to `- [X]`
   - Report progress to user after each task

### 8. Testing Phase (After All Tasks Complete)

**🚨 CRITICAL: ONLY test-writer-fixer WRITES TESTS 🚨**

Implementation agents (frontend-developer, backend-architect, etc.) MUST NOT write tests during task execution. All test creation, execution, and fixing is exclusively handled by test-writer-fixer in this phase.

Once all tasks T{START_TASK} through T{END_TASK} are implemented:

1. **Invoke test-writer-fixer agent** (exclusive test authority):
   ```
   Use Task tool with subagent_type: "test-writer-fixer"
   Prompt: "Test the implementation of tasks T{START_TASK} through T{END_TASK} in spec {SPEC_ID}.
   Review the code changes, run existing tests, write new tests if needed, and fix any failures."
   ```

2. **Iterative testing loop** (max 3 rounds):
   - **Round 1**: test-writer-fixer runs tests and reports results
   - If tests fail:
     - **Round 2**: Invoke appropriate coding agent to fix issues based on test results
     - test-writer-fixer re-runs tests
   - If still failing:
     - **Round 3**: Final iteration with coding agent and test-writer-fixer
   - After 3 rounds, report status to user even if some tests still fail

3. **Validation**:
   - All new tests must pass
   - Existing tests must continue to pass
   - Code quality checks should pass

### 9. Completion Validation (Spec Alignment)

**Cross-reference against specification:**

1. **Read spec.md user stories** related to completed tasks
2. **For each relevant user story**, verify:
   - Success criteria are addressed
   - Acceptance tests would pass
   - Implementation matches specification intent
   - No functionality gaps exist

3. **Report validation results**:
   ```
   📋 Specification Alignment Check:

   User Story 1: "Users can reset password via email"
   ✅ Status: Fully implemented
   - Email sending: ✓
   - Token validation: ✓
   - Password update: ✓

   User Story 2: "Users receive error messages for invalid tokens"
   ⚠️ Status: Partially implemented (missing edge case handling)
   - Token expiry: ✓
   - Invalid format: ✗ (needs implementation)
   - Already used: ✓

   User Story 3: "Password reset follows security best practices"
   ✅ Status: Fully implemented

   Overall Alignment: 2.5/3 stories complete
   ```

4. **Handle gaps**:
   - If gaps exist, document them in session summary
   - Suggest follow-up tasks for incomplete items
   - **Do NOT block completion** (gaps are informational only)
   - Report gaps to user for awareness

5. **Validate implementation quality**:
   - Code follows project patterns (from plan.md)
   - Error handling is comprehensive
   - Edge cases are considered
   - Documentation is adequate

### 10. Session Summary Documentation

**CRITICAL**: Session summary filename format is **NON-NEGOTIABLE**:
- Format: `t###-t###-summary.yaml` (lowercase 't', hyphen-separated)
- Example: `t117-t123.yaml` for tasks T117-T123
- Single task: `t105-t105.yaml` for task T105
- Path: `docs/{SPEC_ID}/session-summary/t{START_TASK}-t{END_TASK}-summary.yaml`

**NEVER create**:
- `README.md`
- `password-reset-flow.yaml`
- Any other filename format

**Invoke code documenter agent**:
```
Use Task tool with subagent_type: "code-documenter"
Prompt: "Document the implementation of tasks T{START_TASK} through T{END_TASK} for spec {SPEC_ID}.

CRITICAL PATH INSTRUCTIONS (STRICT ENFORCEMENT):
- Create session summary at EXACT path: /Users/jbeausoleil/Projects/03_projects/personal/cortex/docs/{SPEC_ID}/session-summary/t{START_TASK}-t{END_TASK}-summary.yaml
- Filename MUST be: t{START_TASK}-t{END_TASK}-summary.yaml (NUMERIC RANGE ONLY, lowercase 't', hyphen-separated)
- CORRECT examples: t117-t123.yaml ✓ | t136-t136.yaml ✓
- INCORRECT examples: t136-transparency-modal.yaml ✗ | t139-audit-triggers.yaml ✗
- DO NOT create README.md or any other filename
- DO NOT create files with descriptive names or feature names in the filename
- DO NOT create ANY other files in docs/{SPEC_ID}/ directory
- DO NOT delegate documentation creation to other agents
- This is the ONLY documentation file permitted for this execution

CRITICAL BREVITY REQUIREMENTS:
- Maximum 400 lines total (target: 250-300 lines)
- Purpose: Scannable reference for end-of-spec triage, NOT comprehensive documentation
- Should take no more than 15 minutes to produce
- Focus: What shipped, key decisions, risks requiring triage

REQUIRED STRUCTURE (Agent-Optimized YAML):

metadata:
  spec_id: \"{SPEC_ID}\"
  task_range: \"T{START_TASK}-T{END_TASK}\"
  date: \"YYYY-MM-DD\"
  branch: \"branch-name\"
  duration: \"X hours\"
  status: \"completed\"

executive_summary:
  description: \"Brief 1-2 sentence summary of what was accomplished\"
  key_achievements:
    - \"Achievement 1 with metric\"
    - \"Achievement 2 with metric\"

tasks_completed:
  - task_id: \"TXXX\"
    description: \"Brief task description\"
    files_modified:
      - \"path/to/file1.ts\"
      - \"path/to/file2.ts\"
    features:
      - \"Feature 1\"
      - \"Feature 2\"

key_decisions:
  - decision: \"Decision name\"
    rationale: \"Brief explanation\"
    impact: \"What this affects\"

test_results:
  summary: \"X/Y tests passing\"
  by_category:
    unit_tests: \"pass/total\"
    integration_tests: \"pass/total\"
  coverage:
    overall: \"X%\"
    critical_paths: \"100%\"

git_commits:
  - sha: \"commit-hash\"
    message: \"Commit message\"
    files_modified:
      - \"path/to/file\"
    stats:
      additions: 123
      deletions: 45

risks_and_backlog:
  high_priority:
    - issue: \"Issue description\"
      impact: \"Impact description\"
      solution: \"Proposed solution\"
      estimate: \"X hours\"
  medium_priority:
    - issue: \"Issue description\"
      impact: \"Impact description\"
      solution: \"Proposed solution\"
      estimate: \"X hours\"
  low_priority:
    - issue: \"Issue description\"
      impact: \"Impact description\"
      solution: \"Proposed solution\"
      estimate: \"X hours\"

next_steps:
  immediate:
    - \"Task before production\"
  next_session:
    - \"TXXX: Task description\"

metrics:
  development:
    files_modified: 10
    lines_added: 500
    lines_deleted: 200
    tests_written: 20
  quality:
    test_coverage: \"80%\"
    compliance_checks: \"5/5 passed\"
  session_performance:
    duration: \"4 hours\"
    quality_rating: \"⭐⭐⭐⭐⭐\"

constitutional_compliance:
  - principle: \"Privacy-First\"
    status: \"✅ Compliant\"
    evidence: \"Brief evidence\"

WHAT TO EXCLUDE:
- ❌ Code examples or snippets (file paths only)
- ❌ Detailed test breakdowns (use counts/status only)
- ❌ Prose paragraphs (use structured data)
- ❌ Redundant information (keep it scannable)"
```

### 10b. Documentation Validation (Immediate Enforcement)

**After code-documenter completes, validate documentation compliance:**

1. **Check permitted file exists:**
   ```bash
   # Verify the session summary was created at the correct path
   if [ -f "docs/{SPEC_ID}/session-summary/t{START_TASK}-t{END_TASK}-summary.yaml" ]; then
     echo "✓ Session summary created correctly"
   else
     echo "✗ ERROR: Session summary missing or wrong path"
   fi
   ```

2. **Scan for violations in docs/{SPEC_ID}/**:
   - Use bash tool to find all documentation files in `docs/{SPEC_ID}/`
   - Identify any file that does NOT match the exact permitted path
   - Common violations:
     - Wrong filename format (descriptive names, feature names)
     - Files outside `session-summary/` directory
     - Additional documentation files (test reports, checklists, guides)

3. **Delete all violations:**
   ```bash
   # Delete any YAML/markdown file that isn't the permitted session summary
   find docs/{SPEC_ID} -type f \( -name "*.yaml" -o -name "*.md" \) ! -path "*/session-summary/t{START_TASK}-t{END_TASK}-summary.yaml" -delete
   ```

4. **Report actions taken:**
   - If violations found: List deleted files with rationale
   - If no violations: Confirm documentation policy compliance
   - Example report:
     ```
     🗑️  Documentation Policy Enforcement:
     Removed 3 unauthorized files:
     - docs/002-auth/audit-event-coverage.yaml (supplementary doc)
     - docs/002-auth/session-summary/t136-transparency-modal.yaml (wrong filename format)
     - docs/002-auth/test-report-t135-t140.yaml (test report)

     ✓ Only permitted file remains: docs/002-auth/session-summary/t135-t140.yaml
     ```

**This validation is MANDATORY and non-negotiable.** Execute immediately after code-documenter completes, before proceeding to Section 11.

### 11. Final Validation and Report

1. **Verify session summary created**:
   - Check file exists at: `docs/{SPEC_ID}/session-summary/t{START_TASK}-t{END_TASK}-summary.yaml`
   - Verify filename matches exact format (lowercase 't', hyphen-separated, .yaml extension)
   - If incorrect format, regenerate with code-documenter

2. **Report completion**:
   ```
   ✅ Tasks T{START_TASK} through T{END_TASK} completed

   📝 Implementation Summary:
   - Tasks executed: {count}
   - Files modified: {list}
   - Tests status: {pass/fail counts}
   - Git commits: {count} commits

   📄 Session Summary: docs/{SPEC_ID}/session-summary/t{START_TASK}-t{END_TASK}-summary.yaml

   Next steps: [Any recommended follow-up work]
   ```

## 12. Agent Selection Guidelines (Detailed Reference)

### Primary Implementation Agents (Choose ONE per task)

These agents are mutually exclusive based on the task's primary technical domain. Each task should be assigned to exactly ONE primary agent:

#### Engineering Domain

**frontend-developer** - UI/Component Implementation
- **Use for:** React components, Next.js pages, client state management, CSS/styling, responsive design, accessibility features
- **Task indicators:** "implement UI", "create component", "add page", "style", "responsive", "accessibility"
- **File patterns:** `src/components/`, `src/app/`, `*.tsx`, `*.css`, `tailwind.config.*`
- **Example tasks:** "Create transaction list component with filtering", "Implement mobile-responsive budget dashboard"

**backend-architect** - Server/API/Database Implementation
- **Use for:** API endpoints, database schemas, migrations, authentication, server logic, data validation, Edge Functions
- **Task indicators:** "API", "endpoint", "database", "migration", "auth", "RLS", "trigger", "function"
- **File patterns:** `supabase/migrations/`, `supabase/functions/`, `src/lib/api/`, `*.sql`
- **Example tasks:** "Create transaction CRUD API", "Implement RLS policies for budget sharing", "Set up database triggers for audit logging"

**ai-engineer** - ML/AI Feature Implementation
- **Use for:** ML models, AI integrations, forecasting algorithms, sentiment analysis, recommendation engines, OpenAI API integration
- **Task indicators:** "ML", "AI", "forecast", "predict", "sentiment", "inference", "model", "OpenAI"
- **File patterns:** `src/lib/ml/`, `src/lib/ai/`, model configs, inference pipelines
- **Example tasks:** "Implement Prophet forecasting model", "Integrate OpenAI API for emotional context detection", "Build recommendation engine for spending insights"

**rapid-prototyper** - New Feature Scaffolding & MVP
- **Use for:** New feature initialization, proof-of-concepts, experimental features, scaffolding multi-file feature structures
- **Task indicators:** "scaffold", "prototype", "MVP", "new feature", "proof-of-concept", "initial setup"
- **File patterns:** Multiple new files across domains
- **Example tasks:** "Scaffold emotional tracking feature structure", "Create MVP for budget scenario simulator"

**devops-automator** - Infrastructure & Deployment
- **Use for:** CI/CD configuration, deployment scripts, environment setup, monitoring, build optimization, infrastructure-as-code
- **Task indicators:** "CI/CD", "deploy", "pipeline", "Docker", "monitoring", "build", "environment"
- **File patterns:** `.github/workflows/`, `Dockerfile`, deployment configs, `vercel.json`
- **Example tasks:** "Set up GitHub Actions for automated testing", "Configure Vercel deployment with environment variables"

**mobile-app-builder** - Mobile-Specific Implementation
- **Use for:** React Native components, mobile navigation, platform-specific features, mobile performance optimization, native modules
- **Task indicators:** "mobile", "React Native", "iOS", "Android", "native", "platform-specific"
- **File patterns:** Mobile-specific directories, platform configs
- **Example tasks:** "Implement mobile transaction entry screen", "Add biometric authentication for mobile"

#### Design Domain

**ui-designer** - Design System & Visual Implementation
- **Use for:** Design systems, component libraries, visual design implementation, brand consistency, design tokens
- **Task indicators:** "design system", "component library", "visual design", "brand", "tokens"
- **File patterns:** `src/design/`, design system configs, Storybook files
- **Example tasks:** "Create design system for transaction categories", "Implement brand color tokens and theming"

**ux-researcher** - UX Validation & User Research
- **Use for:** User flow validation, usability testing integration, user research analysis, journey mapping implementation
- **Task indicators:** "UX validation", "user testing", "user research", "journey map", "usability"
- **File patterns:** UX documentation, research artifacts
- **Example tasks:** "Implement user testing tracking for onboarding flow", "Validate emotional tagging UX with research data"

#### Support Domain (Constitutional Compliance & Specialized Validation)

**legal-compliance-checker** - Privacy & Compliance Implementation
- **Use for:** GDPR/CCPA compliance features, privacy controls, data export/deletion, audit logging, terms of service implementation
- **Task indicators:** "compliance", "privacy", "GDPR", "CCPA", "audit", "data export", "deletion"
- **File patterns:** Privacy-related features, audit logs, compliance docs
- **Example tasks:** "Implement GDPR-compliant data export feature", "Add audit logging for sensitive operations"

### Decision Tree for Agent Selection

Use this decision tree when selecting an agent for a task:

```
1. Does the task involve ML/AI features or algorithms?
   YES → ai-engineer
   NO → Continue

2. Is this scaffolding a new multi-component feature or MVP?
   YES → rapid-prototyper
   NO → Continue

3. Does the task primarily involve server-side logic, APIs, or database?
   YES → backend-architect
   NO → Continue

4. Does the task primarily involve UI components, pages, or styling?
   YES → frontend-developer
   NO → Continue

5. Does the task involve CI/CD, deployment, or infrastructure?
   YES → devops-automator
   NO → Continue

6. Does the task involve mobile-specific native features?
   YES → mobile-app-builder
   NO → Continue

7. Does the task involve design system or visual design implementation?
   YES → ui-designer
   NO → Continue

8. Does the task involve privacy/compliance features?
   YES → legal-compliance-checker
   NO → Continue

9. Does the task involve UX research validation?
   YES → ux-researcher
   NO → frontend-developer (default for unspecified implementation)
```

### Automatic Support Agents (Context Only)

These agents are automatically invoked by speckit.execute and do not need to be manually selected:

- **test-writer-fixer** - 🚨 **EXCLUSIVE TEST AUTHORITY** 🚨 The ONLY agent permitted to write, run, or fix tests. Automatically invoked in Section 8 (Testing Phase) after all implementation tasks complete. NO other agent may create test files or test code.
- **code-documenter** - Automatically invoked in Section 10 (Session Summary Documentation) for final documentation

### Optional Specialized Agents (Invoke When Needed)

These agents can be invoked alongside primary agents for specialized validation or analysis:

- **api-tester** - For comprehensive API contract testing and load testing (if task requires API validation)
- **performance-benchmarker** - For performance profiling and optimization validation (if task requires performance validation)
- **test-results-analyzer** - For analyzing complex test failures or quality metrics (if test failures need deep analysis)

### Conflict Resolution Rules

If a task spans multiple domains (e.g., "Create API endpoint and UI for transaction editing"):

1. **Split into sub-tasks** and assign different agents to each sub-task
2. **Primary domain rule:** Choose the agent for the most complex/critical part
3. **Sequential execution:** Execute backend-architect first (API), then frontend-developer (UI)

Example split:
- T117a: Create transaction editing API endpoint → **backend-architect**
- T117b: Create transaction editing UI form → **frontend-developer**

## Error Handling

- If task execution fails: Report error, ask user whether to continue or stop
- If tests fail after 3 rounds: Report failures, document in session summary, continue
- If git commit fails: Report error, ask user to resolve conflicts
- If session summary path is wrong: Stop and fix before completing

## Notes

- This command executes a focused subset of tasks from the specified task file (default: tasks.md)
- Use `/speckit.implement` to execute all tasks in the feature
- Tasks should be marked as completed `[X]` in {TASK_FILE} as they finish
- Deep, sequential thinking is essential - don't rush through tasks
- Test thoroughly before considering tasks complete
- Session summary documentation is mandatory and must follow exact path format

## Usage Examples

```bash
# Execute task 148 from default tasks.md
/speckit.execute T148

# Execute task 148 from tasks_02.md (with @ prefix)
/speckit.execute @tasks_02.md T148

# Execute task range from tasks_02.md (without @ prefix)
/speckit.execute tasks_02.md T148-T150

# Execute single task from alternate file
/speckit.execute tasks_03.md T200
```
