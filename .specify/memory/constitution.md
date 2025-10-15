<!--
  SYNC IMPACT REPORT
  ==================
  Version Change: 1.0.0 → 2.0.0
  Date: 2025-10-14

  Major Changes:
  - Initial constitution creation for SourceBox project
  - Established 7 core principles (verticalization, speed, local-first, boring tech, open source, developer-first, ship fast)
  - Defined 7 technical constraints (performance, distribution, cost, licensing)
  - Added 5 product philosophy statements
  - Established 7 development practices (TDD, Spec-Kit driven, manual QA)
  - Created 5 UX principles (zero config, transparent progress, graceful failures)
  - Defined 5 business constraints (phased monetization, indie constraints)
  - Added 5 critical legal constraints (independent development protection)
  - Established 4 GTM philosophy statements
  - Listed 8 anti-patterns to explicitly avoid

  Templates Status:
  ✅ plan-template.md - Constitution Check section now enforces speed, verticalization, local-first
  ✅ spec-template.md - User scenarios must align with developer-first, independently testable principles
  ✅ tasks-template.md - Task organization must support TDD-first, Spec-Kit workflow

  Follow-up Actions:
  - None - all principles clearly defined with concrete values

  Version Bump Rationale:
  - MAJOR (1.0.0 → 2.0.0): First complete constitution defining all governance rules from scratch
-->

# SourceBox Constitution

## Mission & Vision

**Mission**: Provide developers with instant access to realistic, verticalized demo data through CLI tools and Docker images.

**Vision**: Make demo data effortless so developers can focus on building instead of seeding databases. Transform the days to weeks developers waste per data provisioning cycle into 30 seconds of automated, production-like data generation.

**Core Problem**: Developers waste days to weeks provisioning realistic demo data for each environment. Faker is too generic ("John Doe", "Acme Corp"), manual SQL is tedious, and production dumps are a security nightmare. Data provisioning can take 30+ minutes for simple apps or an entire week for complex data masking with referential integrity. There's no tool for spinning up databases with realistic, verticalized data in 30 seconds.

**Target Users**: Backend developers, data engineers, Sales Engineers, QA engineers, DevRel engineers at tech companies building data-intensive applications.

**North Star Metric**: Developers using SourceBox weekly (target: 10,000+ by Month 12)

---

## Core Principles

### I. Verticalized > Generic (NON-NEGOTIABLE)

Fintech data MUST look like fintech (loan amounts, credit scores, interest rates, proper payment schedules), healthcare MUST look like healthcare (diagnoses, prescriptions, insurance claims with proper coding), retail MUST look like retail (SKUs, inventory levels, seasonal patterns) — not "John Doe" working at "Acme Corp" with employee ID "123".

**Rationale**: Generic data from Faker.js makes demos look unprofessional to investors and customers. Verticalized data is the core competitive differentiator. If a schema isn't 10x more realistic than Faker, don't ship it.

**Implementation Standard**:
- Data distributions MUST match real-world patterns (e.g., loan amounts follow actual lending distributions, not uniform random)
- Relationships MUST be realistic (e.g., credit scores correlate with loan approval rates)
- Edge cases MUST be included (negative values, outliers, nulls where appropriate)
- Industry terminology MUST be accurate (fintech uses "APR", not "interest percentage")

### II. Speed > Features (NON-NEGOTIABLE)

< 30 seconds database seeding time is NON-NEGOTIABLE. < 10 seconds Docker container spin-up is NON-NEGOTIABLE. If a feature makes seeding slower than 30 seconds, CUT THE FEATURE, not the speed. Developers will abandon tools that waste their time.

**Rationale**: SourceBox competes against 2-3 hours of manual SQL work. Must be 10x faster, not 10% faster. Speed is the #1 reason developers will adopt or abandon this tool.

**Performance Gates**:
- Database seeding: 1,000 records in < 30 seconds (measured on standard hardware)
- Docker image: Ready to query in < 10 seconds from `docker run`
- CLI install: < 5 seconds from package manager
- Any operation that breaks these thresholds MUST be rejected or refactored

### III. Local-First, Cloud Optional

CLI and Docker MUST run entirely on a laptop with zero internet connectivity. No cloud APIs, no authentication, no network calls. Cloud version (Phase 2) is optional convenience, not a requirement. Developers own their data and environment.

**Rationale**: Developers work on planes, in coffee shops with bad wifi, in secure/air-gapped environments. Local-first ensures SourceBox works everywhere. Cloud features are monetization, not core value.

**Architecture Requirements**:
- All data generation MUST be deterministic and reproducible offline
- No telemetry or analytics without explicit opt-in
- Docker images MUST be fully self-contained (no external downloads on first run)
- Cloud version (Phase 2) MUST be additive, not replacing local functionality

### IV. Boring Tech Wins

Go for CLI (single binary, cross-platform, boring), Docker for distribution (standard, boring), PostgreSQL/MySQL for databases (proven, boring). Avoid Rust, WebAssembly, exotic databases UNLESS proven necessary through real performance bottlenecks. Simple > complex.

**Rationale**: The goal is to ship and validate, not to experiment with cutting-edge technology. Boring tech has solved problems, has documentation, has community support. Ship working software in 12 weeks, not chase perfect architecture.

**Tech Stack Constraints**:
- CLI: Go + Cobra framework (standard, well-documented, single binary)
- Databases: PostgreSQL, MySQL (Phase 1), MongoDB (Phase 2 only after validation)
- Data generation: Standard libraries + proven patterns (no AI/ML unless necessary)
- Distribution: npm, homebrew, Docker Hub (developers already use these)

### V. Open Source Forever (NON-NEGOTIABLE)

CLI and Docker images MUST be free forever, MIT license, no proprietary extensions, no paywalls. Cloud version (Phase 2) is optional freemium SaaS. Open source is legal protection, community trust, and viral distribution. This is non-negotiable.

**Rationale**: Open source ensures no employer can claim ownership, enables community contributions, builds trust with developers, and creates viral distribution. Monetization comes from cloud convenience (Phase 2), not from locking down core functionality.

**Research Validation:** Market analysis shows no multi-vertical, open-source, affordable competitor exists. Competitors are either:
- Free but generic (Faker.js: "John Doe" problem)
- Specialized but expensive (Tonic: $3K/year, Gretel: $3.5K/year)
- Open but single-vertical (Synthea: healthcare only)

SourceBox's unique position: Free + Verticalized + Multi-industry = No direct competitor.

**Licensing Requirements**:
- MIT license for all CLI and Docker code (maximum permissiveness)
- No CLA (Contributor License Agreement) required — MIT covers it
- Cloud version source code (Phase 2) MAY be closed source (optional monetization)
- Schema definitions MUST be open source and community-contributed

#### The Synthea Model: Open Source as Competitive Moat

**Insight from Research:** Synthea (open-source healthcare simulator) became the industry standard NOT despite being free, but BECAUSE it was free. Hospitals, researchers, and developers adopted it universally because there was no barrier to experimentation.

**SourceBox Strategy:** Apply Synthea's model across EVERY vertical (fintech, healthcare, retail, logistics, etc.)

**Why This is a Moat:**
- Competitors (Tonic, Gretel, Hazy) are proprietary and expensive ($3K-38K/year)
- Generic tools (Faker.js) are free but unrealistic ("John Doe" problem)
- Specialized tools (Synthea) are free AND realistic but single-vertical only
- **SourceBox is the ONLY:** Free + Realistic + Multi-vertical tool

**Competitive Advantages:**
1. **Viral adoption:** Developers try without budget approval, tell their teams
2. **Community contributions:** Domain experts contribute schemas (Workday consultants, healthcare engineers)
3. **Cannot be replicated:** Competitors can't easily open-source proprietary codebases
4. **Network effects:** More schemas → more users → more contributors → more schemas

**Positioning:** "The Synthea for every industry" - proven model, validated by research

### VI. Developer-First Design

CLI-first (not web UI), Docker-native (one command to run), zero config (works out of the box), transparent progress (show progress bars, record counts, time estimates), fail gracefully (show exact error + how to fix). No unnecessary abstractions, no magic, no surprises.

**Rationale**: SourceBox is built by developers, for developers. Developers value speed, clarity, and control over polish and aesthetics. A boring CLI that works is better than a flashy UI that breaks.

**UX Requirements**:
- Zero configuration default: `sourcebox seed mysql --schema=fintech-loans --records=1000` MUST work without setup
- Progress transparency: Show real-time progress bars, record counts, estimated time remaining
- Explicit errors: If database connection fails, show exact connection string + error message + suggested fix
- No silent failures: Every error MUST have actionable guidance
- Boring design: Functional > flashy (developers value reliability over aesthetics)

### VII. Ship Fast, Validate Early

12-week MVP launch (not 12 months). Start with 3 schemas (fintech, healthcare, retail), not 30. Launch on GitHub, submit to Hacker News, validate before building cloud SaaS. TDD required for core functionality. Manual QA before every release. Spec-Kit driven workflow: `/speckit.specify` → `/speckit.plan` → `/speckit.implement`.

**Rationale**: Unknown if developers will adopt SourceBox. Must validate market demand before investing in cloud infrastructure. Rapid iteration beats perfect planning. Get feedback early and often.

**Shipping Discipline**:
- MVP scope: 3 schemas (fintech, healthcare, retail), 2 databases (MySQL, Postgres), 6 Docker images
- 12-week launch deadline (hard constraint)
- Open source launch first: GitHub, Hacker News, Reddit r/programming
- Validation metrics: 1,000 GitHub stars, 10,000 Docker pulls, 100 daily CLI installs
- Cloud version (Phase 2) ONLY after MVP validation (don't build cloud before proving local demand)

---

## Technical Constraints

### 1. Performance (NON-NEGOTIABLE)

- **Database seeding**: < 30 seconds for Tier 1 schemas, < 2 minutes for Tier 2, < 5 minutes for Tier 3 (measured on 2020 MacBook Pro)
- **Docker spin-up**: < 10 seconds from `docker run` to queryable database
- **CLI install**: < 5 seconds from `npm install -g sourcebox` or `brew install sourcebox`
- **Memory footprint**: < 100MB RAM for CLI, < 200MB for Docker container
- **Disk space**: < 50MB per Docker image (compressed)

**Why**: Speed is the core value proposition. Break these thresholds and developers abandon the tool.

#### Performance Tiers by Schema Complexity

**Tier 1 - Simple Schemas** (<30 seconds, MVP focus)
- 1-5 entities with straightforward relationships
- Examples: fintech loans, retail orders, basic user profiles
- Target: < 30 seconds for 1,000 records
- Use case: Backend developers, quick prototypes

**Tier 2 - Medium Schemas** (<2 minutes, Phase 2)
- 10-50 entities with moderate complexity
- Examples: healthcare patients (with visits/prescriptions), SaaS analytics, e-commerce (orders + inventory + customers)
- Target: < 2 minutes for 1,000 records
- Use case: Sales Engineers, QA environments

**Tier 3 - Complex Schemas** (<5 minutes, Community-driven)
- 100+ entities with deep referential integrity
- Examples: Workday HCM, SAP ERP, enterprise CRM
- Target: < 5 minutes for complete environment setup
- Use case: Implementation consultants, enterprise demos
- **Strategy:** Community-contributed, not MVP scope

**Constitutional Constraint:** The "<30 seconds" requirement applies to **Tier 1 schemas only**. Tier 2/3 schemas may take longer while still providing massive time savings vs manual provisioning (which takes days/weeks).

### 2. Distribution Channels (REQUIRED)

- **CLI**: npm (for Node.js developers), homebrew (for macOS developers), apt/yum (for Linux developers)
- **Docker**: Docker Hub (standard container registry)
- **Documentation**: Docusaurus site (standard for dev tools, SEO-friendly)
- **Source**: GitHub (standard for open source, community trust)

**Why**: Developers already use these channels. Avoid custom distribution mechanisms.

### 3. Database Support (PHASED)

- **Phase 1 (MVP)**: MySQL, PostgreSQL (covers 90% of developer use cases)
- **Phase 2**: MongoDB (NoSQL use cases, only after MVP validation)
- **Future**: SQLite, Redis, Cassandra (based on community demand)

**Why**: Start with proven relational databases. Expand based on validated demand, not speculation.

### 4. Cost Constraints

- **Phase 1 (CLI/Docker)**: $0/user (free forever, MIT license)
- **Phase 2 (Cloud SaaS)**: < $2/user/month infrastructure cost (target 98%+ gross margin)
- **Development**: < $100/month total cost (Vercel free tier, GitHub free tier, Cloudflare free tier)

**Why**: Indie project with no funding. Must be capital-efficient. High gross margin enables sustainability.

### 5. Code Quality Standards

- **TypeScript**: Strict mode enabled (no implicit any, strict null checks)
- **Go**: `go vet`, `golangci-lint` passing with zero warnings
- **Testing**: TDD required for data generation, schema validation, CLI commands, Docker builds
- **Test-after OK**: Schema definitions, documentation, examples (not core logic)
- **Coverage**: > 80% for core data generation and CLI logic (not documentation)

**Why**: Quality prevents bugs that erode trust. TDD ensures testable design. Test-after for non-critical paths balances speed and quality.

### 6. Open Source License (NON-NEGOTIABLE)

- **License**: MIT (maximum permissiveness, no restrictions)
- **No CLA**: Contributor License Agreement not required (MIT is sufficient)
- **No proprietary extensions**: All CLI and Docker code MUST be open source
- **Cloud source (Phase 2)**: MAY be closed source (optional monetization)

**Why**: MIT license protects independent development status, builds trust, enables community contributions, ensures no employer ownership claims.

### 7. Platform Support

- **CLI**: macOS (Intel + Apple Silicon), Linux (x86_64 + ARM64), Windows (x86_64)
- **Docker**: Linux containers (standard, runs on all platforms via Docker Desktop)
- **Manual QA**: Test on Mac, Linux, Windows before EVERY release (no CI-only releases)

**Why**: Developers use diverse platforms. Cross-platform support is table stakes for dev tools.

---

## Product Philosophy

### 1. MVP Mindset: Start Small, Expand via Community

Start with 3 schemas (fintech, healthcare, retail), not 30. Launch and validate demand. Expand to 50+ schemas via community contributions. Every industry is different—let domain experts contribute their schemas, not build everything in-house.

**Approach**:
- Phase 1 (MVP): 3 hand-crafted, high-quality schemas (fintech, healthcare, retail)
- Phase 1 success: 1,000 GitHub stars, 10,000 Docker pulls, 100 daily CLI installs
- Phase 2: Open schema contributions (accept PRs for new verticals)
- Phase 2 goal: 50+ schemas via community, not solo development

### 2. Open Source First, Cloud Optional

CLI and Docker MUST be free forever. Cloud version (Phase 2) is optional freemium SaaS for teams who want hosted databases, API access, and team collaboration. Don't build cloud before validating local demand.

**Monetization Model**:
- Phase 1: $0 (free forever, open source, MIT license)
- Phase 2: Freemium SaaS with **dual monetization** (cloud-hosted databases + API access)
- Phase 2 trigger: After MVP validation (1K stars, 10K pulls, NPS 50+)
- Unit economics: < $2/user/month infra cost (98%+ gross margin target)

**Phase 2 Offerings:**

**A. Hosted Databases** (for persistent demo environments)
- Pre-seeded databases (Postgres, MySQL) on-demand
- Persistent URLs: `demo-db-abc123.sourcebox.dev:5432`
- Team collaboration: Multiple developers share demo environment
- **Use case:** Sales Engineers with multiple demos per week

**B. API Access** (for programmatic data generation)
- Generate realistic data via REST API: `POST /api/generate/fintech-loans`
- Use cases:
  - E2E testing: Generate test data in CI/CD pipelines
  - Synthetic data for ML: Train models with realistic data
  - Load testing: Generate millions of records on-demand
  - Data masking: Replace production data with synthetic equivalents
- **Use case:** QA automation, ML engineers, performance testing

**Research Validation:** Tonic ($299/mo), Gretel ($295/mo) both charge premium for API access. Market validates high-margin opportunity.

### 3. Developer-Centric: CLI-First, Docker-Native

CLI-first (not web UI). Docker-native (one command to run). No unnecessary abstractions. Fits into developer workflow (local dev, CI/CD, demo environments). Developers value speed and control over polish.

**Design Priorities**:
1. Speed (30s seeding time, 10s Docker startup)
2. Simplicity (zero config, works out of the box)
3. Transparency (progress bars, error messages with solutions)
4. Reliability (boring tech, manual QA, TDD)
5. Aesthetics (last priority — functional > flashy)

### 4. Verticalized, Not Generic

Fintech data looks like fintech. Healthcare looks like healthcare. Retail looks like retail. Not "Employee 123" or "John Doe" or "Acme Corp". If data isn't 10x more realistic than Faker, don't ship it.

**Quality Bar**:
- Real-world distributions (loan amounts, credit scores, patient visit frequencies)
- Industry terminology (APR, not "interest percentage"; ICD-10 codes, not "diagnosis name")
- Proper relationships (credit scores correlate with loan approval, medications correlate with diagnoses)
- Edge cases (negative values, outliers, nulls where appropriate)

### 5. Local-First: Offline by Default, Cloud Optional

Runs entirely on laptop, no internet required. No telemetry, no analytics (without explicit opt-in). Cloud version (Phase 2) is optional convenience, not core functionality. Developers own their data.

**Offline Requirements**:
- Data generation MUST be deterministic and reproducible without network
- Docker images MUST be fully self-contained (no downloads on first run)
- CLI MUST work in air-gapped environments
- Cloud features (Phase 2) MUST be additive, not replacing local functionality

---

## Development Practices

### 1. TDD Required for Core Functionality (NON-NEGOTIABLE)

Test-Driven Development (TDD) MUST be used for:
- Data generation logic (realistic distributions, relationships, edge cases)
- Schema validation (ensure schemas match specifications)
- CLI commands (seed, list-schemas, version, help)
- Docker builds (images start and seed correctly)

**Workflow**: Write test → User approval → Test fails → Implement → Test passes → Refactor

**Rationale**: TDD ensures testable design, catches edge cases early, prevents regressions. Core logic quality is non-negotiable.

### 2. Test-After OK for Non-Critical Paths

Test-after (or no tests) acceptable for:
- Schema definitions (declarative YAML/JSON, validated by schema tests)
- Documentation (README, tutorials, examples)
- Examples (demo code, quickstarts)

**Rationale**: Balance quality with speed. Don't test what's not executable logic. Focus TDD on core value: data generation and CLI.

### 3. Manual QA Before Every Release (NON-NEGOTIABLE)

Before EVERY release, manually test on:
- macOS (Intel + Apple Silicon)
- Linux (x86_64)
- Windows (x86_64)

**Manual QA Checklist**:
- CLI install via package manager works
- `sourcebox seed mysql --schema=fintech-loans --records=1000` completes in < 30s
- Docker images start in < 10s and are queryable
- All examples in README execute successfully
- Error messages are clear and actionable

**Rationale**: CI doesn't catch cross-platform issues, UX problems, or real-world usage patterns. Manual QA ensures quality.

### 4. Ship CLI + Docker MVP in 12 Weeks (HARD DEADLINE)

12-week launch deadline is NON-NEGOTIABLE. Scope adjusts to fit timeline, not the reverse. MVP includes:
- CLI tool (Go + Cobra)
- 3 schemas (fintech, healthcare, retail)
- 2 databases (MySQL, Postgres)
- 6 Docker images
- Docusaurus documentation site
- Open source launch (GitHub, Hacker News, Reddit)

**Why 12 weeks**: Unknown if developers will adopt SourceBox. Must validate market demand before investing further. Fast iteration beats perfect planning.

### 5. Open Source Launch First, Then Validate

Launch sequence:
1. GitHub repo (MIT license, README, examples)
2. Submit to Hacker News (validation via upvotes, comments)
3. Post on Reddit r/programming (validation via engagement)
4. Measure: GitHub stars, Docker pulls, CLI installs, NPS

**Phase 2 trigger**: Only build cloud SaaS after MVP validation (1K stars, 10K pulls, NPS 50+). Don't build cloud before proving local demand.

### 6. Spec-Kit Driven Workflow (REQUIRED)

Every feature follows the Spec-Kit workflow:
1. `/speckit.specify` — Define user stories, requirements, success criteria
2. `/speckit.plan` — Research, design, contracts, quickstart
3. `/speckit.implement` — Execute tasks, TDD-first, manual QA

**Rationale**: Spec-Kit enforces clear requirements before coding, prevents scope creep, ensures testable design, maintains documentation quality.

### 7. Indie Project Constraints (10-15 Hours/Week)

SourceBox is an indie project until $5K MRR or strong validation (10K stars, 50K pulls, 1K paying users). Development constraints:
- 10-15 hours/week maximum (outside work hours)
- Personal equipment only (MacBook, personal GitHub account)
- No employer resources (no cloud credits, no proprietary knowledge)
- Solo development until validation (then consider hiring/co-founders)

**Rationale**: Legal protection (independent development), capital efficiency (no burn rate), sustainable pace (avoid burnout).

---

## User Experience Principles

### 1. Speed > Features (Zero Compromise)

< 30 seconds database seeding is NON-NEGOTIABLE. If a feature makes seeding slower, CUT THE FEATURE, not the speed. Progress bars MUST show real-time progress (records/sec, time remaining). Developers will abandon slow tools.

**UX Standard**:
- Show progress bar with % complete, records inserted, time remaining
- Optimize for perceived speed (start showing output within 1 second, even if backend still processing)
- Never show "Loading..." without progress indication

### 2. Zero Config (Works Out of the Box)

`sourcebox seed mysql --schema=fintech-loans --records=1000` MUST work without configuration files, environment variables, or setup steps. Sane defaults for everything (localhost, port 3306, root user, no password). Advanced config optional, not required.

**Default Behavior**:
- Database: localhost (assume local MySQL/Postgres installed)
- Credentials: root user, no password (standard local dev setup)
- Output: Pretty-printed table (human-readable, not JSON)
- Records: 1,000 (reasonable default for demos)

**Advanced Config** (optional flags):
- `--host`, `--port`, `--user`, `--password` (for non-standard setups)
- `--output=json` (for CI/CD pipelines)
- `--records=10000` (for larger datasets)

### 3. Transparent Progress (No Silent Operations)

Show real-time progress for all operations:
- Database seeding: Progress bar with records/sec, % complete, time remaining
- Docker startup: Show container initialization steps (starting MySQL, seeding data, ready to query)
- CLI install: Show download progress (package size, % complete)

**Rationale**: Developers hate black boxes. Transparency builds trust and reduces perceived wait time.

### 4. Fail Gracefully (Actionable Error Messages)

If database connection fails, show:
- **Exact connection string attempted**: `mysql://root@localhost:3306/demo`
- **Exact error message**: `Connection refused on port 3306`
- **Suggested fix**: `Ensure MySQL is running: brew services start mysql`

**Error Message Template**:
```
❌ Error: [What went wrong]
🔍 Details: [Technical error message]
💡 Fix: [Actionable solution]
📚 Docs: [Link to troubleshooting guide]
```

**Rationale**: Vague errors ("Database error") waste developer time. Actionable errors enable self-service debugging.

### 5. Boring CLI (Functional > Flashy)

Functional, fast, reliable > flashy animations, colors, emoji. Developers value speed and clarity over aesthetics. Use emojis sparingly (✅, ❌, 💡 only). No ASCII art, no splash screens, no unnecessary output.

**CLI Output Style**:
- Success: `✅ Seeded MySQL database 'demo' with 1,000 loan records (15.2 seconds)`
- Error: `❌ Connection refused on port 3306`
- Progress: `[====================] 100% (1000/1000 records) - 15.2s`
- Avoid: Excessive colors, animations, ASCII art, marketing copy

---

## Business Constraints

### 1. Phase 1: Free Forever (Non-Negotiable)

CLI and Docker MUST be free forever. MIT license. No paywalls, no feature gates, no "premium" versions of local tools. Open source is legal protection, community trust, and viral distribution.

**Monetization**: Phase 1 has zero revenue. This is validation phase, not monetization phase.

### 2. Phase 2: Freemium SaaS (After Validation)

**Trigger**: Only build cloud version after MVP validation (1K GitHub stars, 10K Docker pulls, NPS 50+)

**Pricing Model**:
- **Free Tier:**
  - CLI/Docker (unlimited, self-hosted)
  - Community schemas (unlimited)
  - API: 10K requests/mo (hobbyist/experimentation)

- **$20/mo Starter:**
  - 1 hosted database (Postgres or MySQL)
  - 10 GB storage
  - API access: 100K requests/mo
  - Persistent demo URLs
  - **Target:** Individual Sales Engineers, freelancers, solo developers

- **$50/mo Developer:**
  - 3 hosted databases
  - API access: 1M requests/mo (for CI/CD integration)
  - Priority support
  - **Target:** QA teams, ML engineers, automation-heavy teams

- **$200/mo Teams:**
  - 10 hosted databases
  - API access: 10M requests/mo
  - 100 GB storage
  - SSO, team collaboration, audit logs
  - **Target:** Sales orgs (10-50 SEs), DevRel teams

**Unit Economics Target**:
- Infrastructure: $2/user/mo (AWS RDS + API gateway)
- Gross margin: 90% on Starter, 95% on Developer/Teams
- API calls are near-zero marginal cost (highest margin product)

**Revenue Model at Scale:**
- 10K free users → 500 paid (5% conversion)
- 300 Starter ($20) + 100 Developer ($50) + 100 Teams ($200) = **$31K MRR**
- 95% margin = **$29K gross profit/mo** = sustainable indie business

### 3. Indie Project (Until $5K MRR or Strong Validation)

SourceBox is an indie project until:
- $5K MRR (monthly recurring revenue from cloud version), OR
- Strong validation (10K GitHub stars, 50K Docker pulls, 1K paying users)

**Development Constraints**:
- 10-15 hours/week maximum (outside work hours)
- Personal equipment only (no employer resources)
- Solo development (no hiring until validation)
- < $100/month operating costs (free tiers for hosting, CI/CD)

**Why**: Legal protection (independent development), capital efficiency (no burn rate), sustainable pace (avoid burnout).

### 4. Unit Economics (Cloud Version - Phase 2)

**Target**: 98%+ gross margin for cloud version

**Cost Breakdown** (per user/month):
- Database hosting: $1.00 (AWS RDS or Supabase)
- Storage: $0.50 (S3 or equivalent)
- Networking: $0.30 (data transfer)
- Auth/API: $0.20 (Supabase or Auth0)
- **Total**: < $2.00/user/month

**Revenue Target**: $20/mo base tier → $18/mo gross profit (90% margin)

**Why**: High gross margin enables sustainability without venture funding. Capital-efficient business model.

### 5. Legal Protection: Independent Development Only

**CRITICAL**: This project is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information.

**Legal Constraints** (See next section for full details):
- Personal equipment only (MacBook, personal GitHub account)
- Outside work hours only (10-15 hours/week)
- No employer resources (no cloud credits, no internal tools, no proprietary knowledge)
- Public information only (no insider knowledge, no competitive research using employer data)
- Open source license (MIT) protects against ownership claims

---

## Legal Constraints (CRITICAL - NON-NEGOTIABLE)

### 1. Independent Development Only

This project is developed **independently** on personal equipment (MacBook, personal GitHub account), outside of work hours (evenings, weekends, 10-15 hours/week), with **no use of employer resources or proprietary information**.

**Requirements**:
- Personal hardware only (no employer laptops, no employer cloud accounts)
- Outside work hours only (no development during work time)
- No employer resources (no cloud credits, no internal tools, no proprietary libraries)
- No proprietary information (no employer data, no insider competitive knowledge)

**Why**: Legal protection against ownership claims. Ensures SourceBox is 100% independently developed.

### 2. No Employer References

All documentation MUST exclude references to current or past employers. Company names MAY be used for illustrative purposes only (e.g., "Fintech companies like Stripe use data pipelines") but NEVER as direct references or endorsements.

**Allowed**: "Fintech companies often struggle with demo data" (generic statement)
**Forbidden**: "At [Employer], we faced this problem with [Product]" (specific reference)

**Why**: Avoids legal claims, maintains independence, ensures no confidential information disclosure.

### 3. Public Information Only

All competitive research, schema designs, and technical decisions MUST be based on publicly available information (public documentation, blog posts, academic papers, open source code). No insider knowledge, no employer data, no confidential information.

**Sources Allowed**:
- Public documentation (Stripe API docs, AWS docs, open source repos)
- Blog posts and tutorials (public content)
- Academic papers (published research)
- Open source code (GitHub, GitLab)

**Sources Forbidden**:
- Employer internal documentation
- Employer proprietary data or algorithms
- Confidential competitive research
- Insider knowledge from current/past employment

**Why**: Legal protection, ensures clean-room development, avoids trade secret violations.

### 4. Open Source Protection (MIT License)

MIT license ensures **no proprietary claims can be made** on the codebase. Anyone (including employers) can use the code, but they cannot claim ownership or restrict others from using it.

**License Requirements**:
- MIT license for ALL CLI and Docker code (non-negotiable)
- No CLA (Contributor License Agreement) required
- Cloud version source code (Phase 2) MAY be closed source (optional monetization)
- Schema definitions MUST be open source (community contributions)

**Why**: MIT license is maximum permissiveness. Protects against ownership disputes. Enables community contributions.

### 5. Illustrative Examples Only (No Affiliations)

References to companies or products in documentation are **for illustrative purposes only**. SourceBox is a standalone project with **no affiliation to any third party**.

**Example Usage**:
- ✅ "Fintech applications (like loan management systems) benefit from realistic credit score data"
- ✅ "Healthcare applications (like patient record systems) need realistic diagnoses and prescriptions"
- ❌ "SourceBox is used by [Company]" (unless true and verified)
- ❌ "Built in partnership with [Company]" (unless true and contractually agreed)

**Why**: Avoids false endorsement claims, maintains independence, ensures no misleading marketing.

---

## Go-To-Market Philosophy

### 1. Open Source First: Launch on GitHub, Hacker News, Reddit

Launch sequence:
1. **GitHub repo**: MIT license, polished README, working examples, clear quickstart
2. **Hacker News submission**: "Show HN: SourceBox — Verticalized demo data for developers"
3. **Reddit r/programming**: Post with clear value prop and demo video/GIF
4. **Measure validation**: GitHub stars, Docker pulls, CLI installs, comments, NPS

**Success Metrics** (MVP - Month 6):
- 1,000 GitHub stars
- 10,000 Docker Hub pulls
- 100 daily CLI installs
- NPS 50+ (strong product-market fit)

**Why**: Open source launch is zero-cost distribution. Hacker News and Reddit are high-traffic developer communities. Validation comes from engagement, not spending.

### 2. Product-Led Growth: Free CLI/Docker → Viral Adoption → Paid Cloud (Phase 2)

**Phase 1** (Months 1-6): Free forever CLI/Docker
**Phase 2** (Months 7-12): Freemium cloud SaaS ($0 → $20/mo → $200/mo)

**Conversion Funnel**:
1. Developer finds SourceBox via GitHub, Hacker News, Reddit, or SEO
2. Installs CLI or runs Docker image (free forever)
3. Uses for local development, demos, testing (free forever)
4. Needs cloud-hosted databases for team collaboration (converts to $20/mo)
5. Scales to enterprise usage (converts to $200/mo)

**Target Conversion**: 5% free → paid (industry standard for dev tools)

**Why**: Product-led growth has lower CAC (customer acquisition cost) than sales-led. Free tier drives adoption, paid tier monetizes value.

### 3. Developer Content: SEO Blog Posts, YouTube Tutorials, GitHub Examples

**Content Strategy**:
- SEO blog posts: "How to seed realistic fintech data in 30 seconds", "Verticalized demo data vs Faker.js"
- YouTube tutorials: "Building a fintech demo with SourceBox", "Docker + SourceBox quickstart"
- GitHub examples: Example repos for common use cases (Next.js + SourceBox, FastAPI + SourceBox)

**Distribution Channels**:
- Personal blog (SEO-optimized, Docusaurus site)
- YouTube (dev tool tutorials)
- GitHub (example repos linked from README)
- Dev.to, Hashnode (cross-post blog content)

**Why**: Developer content is zero-cost marketing. SEO drives long-tail traffic. Tutorials build trust and demonstrate value.

### 4. Community-Driven: Schema Contributions, Feature Requests, Bug Reports

**Community Engagement**:
- Accept schema contributions via GitHub PRs (expand from 3 to 50+ schemas)
- Feature requests via GitHub issues (prioritize by upvotes)
- Bug reports via GitHub issues (fix high-impact bugs first)

**Community Goals** (Phase 2 - Month 12):
- 50+ schemas (via community contributions, not solo development)
- 100+ GitHub contributors
- Active Discord/Slack community (if demand exists)

**Why**: Community-driven development scales beyond solo capacity. Schema contributions reduce development burden. Active community signals product-market fit.

---

## Anti-Patterns (What We Say NO To)

### 1. Feature Bloat (Must Solve Real Developer Pain)

❌ Reject features that are "nice-to-have" but don't solve the core problem (realistic demo data in 30 seconds).

**Examples of Rejected Features**:
- Visual schema builder (nice-to-have, adds complexity, delays MVP)
- AI-generated schemas (interesting, but adds cost and complexity)
- Custom data transformations (edge case, not core value prop)

**Decision Rule**: If a feature doesn't make "realistic demo data in 30 seconds" 10x better, reject it.

### 2. Enterprise-First (Must Work in 30 Seconds, Not 6 Weeks)

❌ Reject features that require long implementation cycles (SSO, RBAC, audit logs) until AFTER MVP validation.

**Phase 1** (MVP): CLI + Docker, < 30s seeding, zero config
**Phase 2** (Freemium): Cloud hosting, API access
**Phase 3** (Enterprise): SSO, RBAC, audit logs, air-gapped deployment

**Decision Rule**: Enterprise features come AFTER product-market fit (1K stars, 10K pulls, NPS 50+), not before.

### 3. Complex Pricing in Phase 1 (CLI/Docker Must Be Free Forever)

❌ Reject any paywalls, feature gates, or "premium" versions of CLI/Docker tools.

**Free Forever**: CLI, Docker images, local development
**Paid** (Phase 2 only): Cloud-hosted databases, API access, team collaboration

**Decision Rule**: If it runs locally, it's free forever. Cloud convenience (Phase 2) is the monetization layer.

### 4. Shiny Tech (Boring Tech Wins Unless Proven Necessary)

❌ Reject Rust, WebAssembly, exotic databases, cutting-edge frameworks UNLESS proven necessary through real performance bottlenecks.

**Boring Tech Stack**:
- Go (CLI) — single binary, cross-platform, proven
- Docker (distribution) — standard for developers
- PostgreSQL/MySQL (databases) — proven, widely used
- Docusaurus (documentation) — standard for dev tools

**Decision Rule**: Use proven tech. Only adopt new tech if measurable performance/UX benefit (10x, not 10%).

### 5. Over-Engineering (Simple > Complex, Ship MVP Before Cloud SaaS)

❌ Reject architectural complexity that delays MVP launch. Ship CLI + Docker in 12 weeks, THEN iterate based on feedback.

**MVP Approach**:
- Start with simple file-based schema definitions (YAML/JSON), not schema DSL
- Start with hardcoded data distributions, not ML models
- Start with 3 schemas (fintech, healthcare, retail), not 30

**Decision Rule**: Ship working software in 12 weeks. Perfect architecture comes from real-world feedback, not upfront planning.

### 6. Generic Data (Faker.js Approach Not Acceptable)

❌ Reject generic data ("John Doe", "Acme Corp", "Employee 123"). Fintech MUST look like fintech, healthcare MUST look like healthcare.

**Quality Bar**:
- Real-world distributions (loan amounts follow actual lending patterns, not uniform random)
- Industry terminology (APR, not "interest percentage"; ICD-10 codes, not "diagnosis name")
- Proper relationships (credit scores correlate with loan approval, medications correlate with diagnoses)
- Edge cases (negative values, outliers, nulls where appropriate)

**Decision Rule**: If data isn't 10x more realistic than Faker, don't ship it.

### 7. Premature Optimization (Ship 3 Schemas, Validate, Then Expand)

❌ Reject building 50+ schemas before validation. Start with 3 high-quality schemas (fintech, healthcare, retail), validate demand, THEN expand via community contributions.

**MVP Scope**: 3 schemas (fintech, healthcare, retail)
**Phase 2 Goal**: 50+ schemas (via community contributions)

**Decision Rule**: Don't scale before validation. Prove demand with 3 schemas, then expand.

### 8. Cloud-First (Must Work Offline, No Internet Required)

❌ Reject cloud-first architecture. CLI and Docker MUST work offline, in air-gapped environments, with no internet connectivity.

**Local-First Requirements**:
- Data generation MUST be deterministic and reproducible without network
- Docker images MUST be fully self-contained (no downloads on first run)
- No telemetry or analytics without explicit opt-in

**Decision Rule**: If a feature requires internet, it's Phase 2 (cloud SaaS), not Phase 1 (MVP).

---

## Decision Framework

### How to Use This Constitution

When making decisions (features, architecture, process, business), ask:

1. **Does this align with core principles?**
   - Is it verticalized (not generic)?
   - Is it fast (< 30s seeding)?
   - Is it local-first (works offline)?
   - Is it boring tech (proven, not experimental)?
   - Is it open source (MIT license)?
   - Is it developer-first (CLI, Docker, zero config)?
   - Does it ship fast (12-week MVP)?

2. **Does it violate technical constraints?**
   - Does it break < 30s seeding time?
   - Does it require non-standard distribution channels?
   - Does it require databases not supported in Phase 1 (MySQL, Postgres)?
   - Does it exceed cost constraints (< $100/month in Phase 1)?

3. **Does it follow development practices?**
   - Did we use TDD for core functionality?
   - Did we manually test on Mac, Linux, Windows?
   - Did we follow Spec-Kit workflow (specify → plan → implement)?

4. **Does it protect legal independence?**
   - Is it developed on personal equipment?
   - Is it outside work hours?
   - Does it use only public information?
   - Does it avoid employer references?

5. **Does it avoid anti-patterns?**
   - Does it add feature bloat?
   - Does it require 6-week implementations?
   - Does it add pricing complexity in Phase 1?
   - Does it use shiny tech without proven necessity?
   - Does it create generic (not verticalized) data?

**Decision Rule**: If the answer to ANY question is "no" or uncertain, STOP and re-evaluate the decision against the constitution.

### Conflict Resolution

If principles conflict (e.g., "ship fast" vs "TDD required"), prioritize:

1. **Legal constraints** (non-negotiable, protects the entire project)
2. **Core principles** (defines what SourceBox is)
3. **Technical constraints** (ensures quality and performance)
4. **Development practices** (ensures sustainable process)

**Example Conflict**: "Ship fast" (12-week MVP) vs "TDD required" (takes more time)

**Resolution**: TDD is NON-NEGOTIABLE for core functionality (data generation, CLI commands). Test-after is acceptable for non-critical paths (documentation, examples). Adjust scope to fit 12-week timeline, not skip TDD.

### Amendment Process

This constitution can be amended when:
- New information changes core assumptions (e.g., different tech stack proves 10x faster)
- Market validation requires strategy shift (e.g., Phase 2 cloud demand is lower than expected)
- Legal requirements change (e.g., new open source license benefits)

**Amendment Requirements**:
1. Document proposed change in GitHub issue
2. Justify why change is necessary (data, user feedback, legal advice)
3. Update constitution version (MAJOR = breaking change, MINOR = new principle, PATCH = clarification)
4. Update dependent templates (plan-template.md, spec-template.md, tasks-template.md)
5. Communicate change to all stakeholders (if team grows beyond solo development)

---

## Governance

### Constitution Supremacy

This constitution supersedes all other practices, processes, and decisions. When in doubt, refer to the constitution. If the constitution is unclear, propose an amendment.

### Compliance Verification

All feature specifications, implementation plans, and pull requests MUST verify compliance with:
- Core principles (7 principles)
- Technical constraints (7 constraints)
- Development practices (7 practices)
- Legal constraints (5 constraints, CRITICAL)

**Verification Checklist** (required for every spec/plan):
- [ ] Aligns with core principles (verticalized, fast, local-first, boring tech, open source, developer-first, ship fast)
- [ ] Meets technical constraints (performance, distribution, database support, cost, quality, license, platform support)
- [ ] Follows development practices (TDD, manual QA, Spec-Kit workflow, 12-week timeline, indie constraints)
- [ ] Protects legal independence (personal equipment, outside work hours, no employer resources, public info only, no references)
- [ ] Avoids anti-patterns (feature bloat, enterprise-first, complex pricing, shiny tech, over-engineering, generic data, premature optimization, cloud-first)

### Complexity Justification

Any decision that violates a constitutional principle MUST be explicitly justified:
- Why is the violation necessary?
- What simpler alternative was rejected, and why?
- What is the measured benefit (10x improvement, not 10%)?

**Example**: Using Rust instead of Go for CLI (violates "boring tech wins")
**Justification Required**: Rust provides [measurable performance benefit], Go approach was rejected because [specific limitation], benchmark shows [10x speedup].

### Runtime Development Guidance

For day-to-day development guidance beyond constitutional principles, refer to:
- **README.md**: Project overview, quickstart, examples
- **CONTRIBUTING.md**: Development workflow, PR guidelines, coding standards
- **docs/architecture.md**: Technical architecture, design decisions
- **Spec-Kit templates**: Feature specification, implementation planning, task generation

---

**Version**: 2.0.0
**Ratified**: 2025-01-14
**Last Amended**: 2025-01-14

---

**How to Use This Constitution**

1. **Before starting a feature**: Read relevant principles (Core Principles, Product Philosophy, Anti-Patterns)
2. **During planning**: Verify compliance with Technical Constraints and Development Practices
3. **During implementation**: Follow TDD discipline, manual QA requirements, Spec-Kit workflow
4. **Before release**: Verify Legal Constraints (no employer references, public info only, MIT license)
5. **When in doubt**: Refer to Decision Framework or propose a constitution amendment

**This constitution is a living document. Amend it when reality changes, but always document the rationale.**
