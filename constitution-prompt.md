# SourceBox Constitution Prompt

> **Usage**: Feed this entire prompt to `/speckit.constitution` to generate the project constitution

---

## Project Context

**Project Name**: SourceBox

**Mission**: Provide developers with instant access to realistic, verticalized demo data through CLI tools and Docker images.

**Vision**: Make demo data effortless so developers can focus on building instead of seeding databases.

**Core Problem Being Solved**: Developers waste 5-10 hours per week creating realistic demo data. Faker is too generic ("John Doe", "Acme Corp"), manual SQL is tedious, and production dumps are a security nightmare. There's no tool for spinning up databases with realistic, verticalized data in 30 seconds.

**Target Users**: Backend developers, data engineers, Sales Engineers, QA engineers, DevRel engineers at tech companies building data-intensive applications.

**Key Insight**: Developers can't demo credibly with generic Faker data ("Employee 123" looks unprofessional to investors/customers). They need verticalized data - fintech data should look like fintech (loan amounts, credit scores), healthcare should look like healthcare (diagnoses, prescriptions, insurance claims).

---

## Constitution Generation Instructions

Please generate a project constitution for SourceBox that includes:

### 1. Core Principles (5-7 principles)
- What are the non-negotiable values that will guide all decisions?
- Think about: speed, simplicity, developer experience, open source, verticalization, local-first
- Key themes: "Ship fast, validate early", "Boring tech wins", "Developer-first design", "Open source forever", "Verticalized not generic", "Local-first, cloud optional"

### 2. Technical Constraints (5-7 constraints)
- What technical decisions are non-negotiable?
- Think about: TypeScript strict mode, Go for CLI, PostgreSQL/MySQL support, < 30s seed time
- Distribution: npm, homebrew, Docker Hub (standard developer channels)
- Performance: < 30s database seeding, < 10s Docker container spin-up
- Cost: < $2/user/month for cloud version (Phase 2), free forever for CLI/Docker
- Open Source: MIT license, no proprietary extensions

### 3. Product Philosophy (3-5 statements)
- How do we approach building features?
- MVP mindset: Start with 3 schemas (fintech, healthcare, retail), expand to 50+
- Open source first: CLI + Docker must be free forever, cloud version (Phase 2) is optional
- Developer-centric: CLI-first, Docker-native, no unnecessary abstractions
- Verticalized: Fintech data looks like fintech, healthcare looks like healthcare (not "Employee 123")
- Local-first: Runs entirely on laptop, no cloud required (cloud optional in Phase 2)

### 4. Development Practices (5-7 practices)
- How do we write code and ship features?
- TDD required for: data generation, schema validation, CLI commands, Docker builds
- Test-after OK for: schema definitions, documentation, examples
- Ship CLI + Docker MVP in 12 weeks (not 12 months)
- Open source launch on GitHub, Hacker News, Reddit (validation before building cloud version)
- Manual QA before every release (test on Mac, Linux, Windows)
- Spec-Kit driven: Always start with `/speckit.specify` → `/speckit.plan` → `/speckit.implement`

### 5. User Experience Principles (3-5 principles)
- How do we design for developers?
- Speed > features (30s seed time non-negotiable, cut features to hit speed)
- Zero config (works out of the box: `sourcebox seed mysql --schema=fintech-loans --records=1000`)
- Transparent progress (show progress bars, record counts, time estimates)
- Fail gracefully (if DB connection fails, show exact connection string + error)
- Boring CLI (functional, fast, reliable - not flashy - developers value speed over design)

### 6. Business Constraints (3-5 constraints)
- What business decisions are non-negotiable?
- Phase 1: Free forever (open source, MIT license, no paywalls in CLI/Docker)
- Phase 2: Freemium SaaS ($0 → $20/mo → $200/mo for cloud-hosted databases)
- Unit economics: 98%+ gross margin for cloud version (< $2/user/mo infra cost)
- Indie project: 10-15 hours/week until $5K MRR or strong validation
- Legal: Independent development only (personal equipment, outside work hours, no employer resources)

### 7. Go-To-Market Philosophy (3-5 statements)
- How do we acquire and retain users?
- Open source first: Launch on GitHub, submit to Hacker News, post on Reddit r/programming
- Product-led growth: Free CLI/Docker → viral adoption → convert to paid cloud in Phase 2
- Developer content: SEO blog posts, YouTube tutorials, GitHub examples
- Community-driven: Accept schema contributions, feature requests, bug reports from users

### 8. Legal Constraints (CRITICAL)
- **Independent Development Only**: This project is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information.
- **No Employer References**: All documentation must exclude references to current or past employers.
- **Public Information Only**: All competitive research, schema designs, and technical decisions must be based on publicly available information.
- **Open Source Protection**: MIT license ensures no proprietary claims can be made on the codebase.
- **Illustrative Examples Only**: References to companies or products in documentation are for illustrative purposes only. SourceBox is a standalone project with no affiliation to any third party.

### 9. What We Say NO To (5-7 anti-patterns)
- What will we explicitly avoid?
- Feature bloat (must solve real developer pain, not nice-to-have)
- Enterprise-first (no 6-week implementations, must work in 30 seconds)
- Complex pricing in Phase 1 (CLI/Docker must be free forever, no paywalls)
- Shiny tech (boring tech wins: Go, Docker, PostgreSQL/MySQL - not Rust/WebAssembly unless proven necessary)
- Over-engineering (simple > complex, ship CLI MVP before building cloud SaaS)
- Generic data (Faker.js approach - "John Doe" is not acceptable, must be verticalized)
- Premature optimization (ship 3 schemas, validate, then expand to 50+)
- Cloud-first (must work offline, no internet required for CLI/Docker)

---

## Key Context for Constitution Generation

### User Persona (Primary)
**Alex the Backend Developer at SaaS Startup**:
- Frustrated: Spends 5-8 hours/week creating demo data with Faker or manual SQL
- Technical: Full-stack developer, knows TypeScript/Python/SQL, builds fintech analytics
- Goal: Ship realistic demo to investors, get seed funding, validate product
- Quote: "I need a database with 1,000 realistic loan records in 30 seconds, not 'John Doe' garbage"

### User Persona (Secondary)
**Maria the Data Engineer**:
- Frustrated: Can't use production data (PII risk), Faker is too generic for testing ETL pipelines
- Technical: Expert in SQL, data pipelines, dbt, Airflow
- Goal: Test ETL pipeline with realistic healthcare data (diagnoses, prescriptions, insurance claims)
- Quote: "I need realistic healthcare data with proper relationships, not random values"

### User Persona (Tertiary)
**Sam the Sales Engineer**:
- Frustrated: Shared demo environments have wrong data, can't spin up custom demos per customer
- Technical: Former developer, knows Docker, SQL, can run CLI commands
- Goal: Spin up custom demo database for healthcare customer in 30 seconds
- Quote: "I want `docker run sourcebox/mysql-healthcare` and have it just work"

### Product North Star
**Developers using SourceBox weekly** (target: 10,000+ by Month 12)

### Competitive Differentiation
- **Only tool** that provides verticalized demo data (fintech looks like fintech, not "Acme Corp")
- **10x faster** (30s vs 2-3 hours for manual SQL, vs web UI for Mockaroo)
- **Free forever** (open source vs $60/mo for Mockaroo, vs manual effort)
- **Developer-native** (CLI + Docker, not web UI - fits developer workflow)
- **Unsolved problem** (no existing tool does this - Faker is generic, Mockaroo is web-based, manual SQL is tedious)

### MVP Scope (12 Weeks)
- **CLI Tool**: Go + Cobra framework, commands: `seed`, `list-schemas`, `version`, `help`
- **3 Verticalized Schemas**: Fintech (loans, borrowers, payments, credit_scores), Healthcare (patients, visits, prescriptions, insurance_claims), Retail (products, orders, customers, inventory)
- **2 Databases**: MySQL, PostgreSQL (MongoDB in Phase 2)
- **6 Docker Images**: `sourcebox/mysql-fintech`, `mysql-healthcare`, `mysql-retail`, `postgres-fintech`, `postgres-healthcare`, `postgres-retail`
- **Documentation**: Docusaurus site with examples, CLI reference, schema library
- **Open Source Launch**: GitHub repo, MIT license, Hacker News post, Reddit post

### Success Metrics (MVP - Month 6)
- 1,000 GitHub stars
- 10,000 Docker Hub pulls
- 100 daily CLI installs (npm + homebrew)
- NPS 50+ (strong product/market fit)

### Success Metrics (Phase 2 - Month 12)
- $10K MRR (cloud version: freemium SaaS)
- 1,000 paying users
- 5% free → paid conversion

### Key Risks to Address in Constitution
1. **Market validation** (will developers actually use this?) → Open source launch first, validate before building cloud SaaS
2. **Schema complexity** (every industry is different) → Start with 3 schemas, expand via community contributions
3. **Data realism** (how realistic is "realistic"?) → Proper distributions, relationships, edge cases (not just random Faker)
4. **Switching costs** (developers have existing workflows) → Must be 10x better (30s vs 2-3 hours), not 10%
5. **Legal/IP** (employment constraints) → Independent development only, personal equipment, outside work hours, no employer resources

---

## Tone and Style for Constitution

The constitution should be:
- ✅ **Opinionated**: Take strong stances (verticalized > generic, open source > proprietary, local-first > cloud-only)
- ✅ **Practical**: Focus on shipping and learning, not perfection
- ✅ **Developer-centric**: Written for/by developers who understand the pain
- ✅ **Concise**: Short, memorable statements (not essays)
- ✅ **Actionable**: Clear guidance for decision-making
- ✅ **Legal-aware**: Protect independent development status in all decisions

Avoid:
- ❌ Corporate jargon ("synergy", "best-in-class", "innovative")
- ❌ Vague statements ("We believe in quality" - everyone does)
- ❌ Contradictions (can't be "move fast" AND "perfect code")
- ❌ Employer references (no Fivetran, no proprietary info)

---

## Example Principle Format (for inspiration)

**Good Example**:
> "Verticalized > Generic. Fintech data must look like fintech (loan amounts, credit scores, interest rates), not 'John Doe' working at 'Acme Corp'. If a schema isn't 10x more realistic than Faker, don't ship it."

**Bad Example**:
> "We value realistic data and strive to deliver the best possible developer experience through innovative data generation solutions."

---

## Suggested Structure for Generated Constitution

```
# SourceBox Project Constitution

## Mission & Vision
[1-2 paragraphs]

## Core Principles
1. [Principle Name]: [Description + rationale]
2. ...

## Technical Constraints
1. [Constraint Name]: [What + why]
2. ...

## Product Philosophy
[3-5 statements about how we build]

## Development Practices
[5-7 specific practices]

## User Experience Principles
[3-5 UX guidelines]

## Business Constraints
[3-5 business rules]

## Legal Constraints (CRITICAL)
[5 legal protection rules]

## Go-To-Market Philosophy
[3-5 GTM principles]

## Anti-Patterns (What We Say NO To)
[5-7 things we explicitly avoid]

## Decision Framework
[How to use this constitution when making decisions]

---
Last Updated: [Date]
Version: 2.0 (Developer Platform Vision)
```

---

## Additional Notes

**Philosophy Inspirations**:
- Docker: "Build, ship, run anywhere" (developer-centric)
- Stripe: "Increase the GDP of the internet" (clear mission)
- Linux: "Free as in freedom" (open source ethos)
- Homebrew: "The missing package manager" (fill obvious gap)

**For SourceBox**:
- "Verticalized demo data in 30 seconds" (clear value prop)
- "Local-first, cloud optional" (developer control)
- "Built by developers, for developers" (founder/market fit)
- "Free forever" (open source commitment)

**Key Tensions to Resolve**:
1. Speed vs realism → Fast is critical (30s), but data must be 10x better than Faker
2. Simple vs powerful → Start simple (3 schemas), expand via community (50+ schemas)
3. Free vs monetization → CLI/Docker free forever, cloud SaaS optional (Phase 2)
4. Local vs cloud → Local-first (works offline), cloud optional (convenience)
5. Indie vs venture → Bootstrap until $10K MRR, then decide on funding
6. Legal vs growth → Independent development only, no employer resources/info

---

## Legal Protection Requirements

**All constitution principles must support**:
1. Independent development on personal equipment
2. Outside work hours only (10-15 hours/week)
3. No use of employer resources or proprietary information
4. No references to current/past employers in documentation
5. Public information only (no insider knowledge)
6. Open source license (MIT) for legal protection
7. Illustrative examples only (no endorsements or affiliations)

**If a decision violates these constraints, it must be rejected.**

---

**End of Constitution Prompt**

Please generate a comprehensive constitution for SourceBox following the guidelines above.
