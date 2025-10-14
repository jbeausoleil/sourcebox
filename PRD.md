# SourceBox: Product Requirements Document (PRD)

**Version**: 2.0 (Developer Platform Edition)
**Last Updated**: 2025-01-14
**Status**: Pre-Launch / MVP Planning
**Owner**: Founder / Product Lead

---

## Executive Summary

**Product Name**: SourceBox

**One-Liner**: Realistic, verticalized demo data for developers - spin up databases with fintech, healthcare, or retail data in 30 seconds.

**Problem**: Developers waste 5-10 hours per week creating realistic demo data. Faker is too generic ("John Doe", "Acme Corp"), manual SQL is tedious, and production dumps are a security nightmare.

**Solution**: SourceBox provides **verticalized demo data** via CLI tool + Docker images. Fintech data looks like fintech (loan amounts, credit scores), healthcare looks like healthcare (diagnoses, prescriptions), retail looks like retail (products, orders).

**Target Users**: Backend developers, data engineers, sales engineers, DevRel, QA engineers

**Success Metric**: Save developers 5+ hours per week on demo data creation

---

## Table of Contents

1. [Problem Statement](#problem-statement)
2. [User Personas](#user-personas)
3. [User Stories](#user-stories)
4. [Product Vision](#product-vision)
5. [MVP Scope](#mvp-scope)
6. [Feature Requirements](#feature-requirements)
7. [Non-Functional Requirements](#non-functional-requirements)
8. [User Experience](#user-experience)
9. [Success Metrics](#success-metrics)
10. [Competitive Analysis](#competitive-analysis)
11. [Go-To-Market](#go-to-market)
12. [Roadmap](#roadmap)
13. [Dependencies & Risks](#dependencies--risks)

---

## Problem Statement

### The Core Problem

**Developers waste 5-10 hours per week creating realistic demo data for:**
- Building features (backend APIs, analytics dashboards)
- Testing data pipelines (ETL, transformations)
- Customer demos (sales engineering)
- QA testing (edge cases, data scenarios)
- Tutorials & documentation (blog posts, video guides)

**Why Current Solutions Fail:**

| Solution | Problem |
|----------|---------|
| **Faker.js** | Too generic ("John Doe", "test@test.com") - not industry-specific |
| **Manual SQL** | Tedious to write 1,000+ INSERT statements with relationships |
| **Production dumps** | Security risk (PII), GDPR/HIPAA nightmare, data drift |
| **Copy/paste tutorials** | Outdated, incomplete, doesn't fit your use case |

**The Missing Piece**: **Verticalized, realistic demo data** that looks like real industries:
- **Fintech**: Loan amounts follow realistic distributions, credit scores are bell-curved around 680
- **Healthcare**: Diagnoses use real ICD-10 codes, medications are actual drugs (Lipitor, Metformin)
- **Retail**: Products are real (iPhone 15 Pro, Nike Air Max), prices make sense

### Quantified Impact

| Metric | Current State | Cost/Impact |
|--------|---------------|-------------|
| **Time spent per week on demo data** | 5-10 hours | $250-500/week (at $50/hr) |
| **Developers affected globally** | 1M+ backend devs, 100K+ data engineers | $1B+ wasted time annually |
| **Pain points** | Generic data → demos don't resonate, lost credibility | Lower conversion rates |

### Why This Matters

**Impact on Developers**:
- ⏰ Time waste (5-10 hours/week that could be building features)
- 😡 Frustration (writing INSERT statements for hours)
- 😰 Demo anxiety ("Will this data look realistic enough?")

**Impact on Business**:
- ❌ Lower demo quality (generic data doesn't resonate with prospects)
- ❌ Slower feature development (waiting for data setup)
- ❌ Security risks (using production dumps with PII)

---

## User Personas

### Primary Persona: Alex the Backend Developer

**Demographics**:
- Age: 28
- Role: Backend Developer at SaaS startup (3 years experience)
- Location: Remote (US)
- Salary: $120K base
- Tech Stack: Next.js, Prisma, PostgreSQL, TypeScript

**Background**:
- Building a SaaS analytics dashboard for fintech companies
- Needs realistic loan data to test API endpoints and show investors
- Currently using Faker.js (generates "John Doe" and "Acme Corp")

**Goals**:
- ✅ Build features fast (no time wasted on data)
- ✅ Demo to investors with realistic data
- ✅ Test edge cases (negative credit scores, large loan amounts)
- ✅ Ship product in 6 weeks

**Frustrations**:
- 😡 "Faker gives me 'John Doe' - investors laugh when I demo"
- ⏰ "Spent 6 hours writing SQL INSERTs for 1,000 loan records"
- 😰 "My demo data doesn't look like real fintech data"

**Day in the Life**:
- **9 AM**: Start building analytics dashboard API
- **10 AM**: Need to test with realistic data → searches "fintech demo data"
- **11 AM**: Finds Faker.js → generates "Employee 123" and "Company ABC"
- **12 PM**: Frustrated, starts writing manual SQL INSERTs
- **3 PM**: Still writing SQL (500 rows done, 500 to go)
- **5 PM**: Finally has data, but it's generic and not realistic
- **6 PM**: Demo to investors → "Why is everything named 'Test Company'?" → Lost credibility

**What Alex Wants**:
> "Just give me a database with 1,000 realistic fintech loans in 30 seconds. I don't have time to write SQL all day."

---

### Secondary Persona: Maria the Data Engineer

**Demographics**:
- Age: 32
- Role: Data Engineer at mid-sized company
- Salary: $140K base
- Tech Stack: Python, SQL, dbt, Airflow, Snowflake

**Background**:
- Builds ETL pipelines for customer data warehouses
- Needs to test transformations with realistic source data
- Currently uses production dumps (security risk) or writes manual SQL

**Goals**:
- ✅ Test data pipelines with edge cases (missing fields, NULL values, outliers)
- ✅ Demo pipeline to customers with their industry data (healthcare, retail)
- ✅ Avoid using production data (GDPR/HIPAA compliance)

**Frustrations**:
- 😡 "Can't use production data (PII risk) but Faker is too generic"
- ⏰ "Spent 8 hours creating test data for a single pipeline demo"
- 😰 "Customer asked 'Is this healthcare data?' → No, it's fake generic data"

**What Maria Wants**:
> "Give me realistic healthcare data with diagnoses, prescriptions, and insurance claims. I need to test my pipeline and demo to customers."

---

### Tertiary Persona: Sam the Sales Engineer

**Demographics**:
- Age: 30
- Role: Sales Engineer at data integration company
- Salary: $150K base + $30K variable
- Demos: 3-4 per week to enterprise customers

**Background**:
- Demos data pipelines to customers (source → transformation → warehouse)
- Needs realistic source data that matches customer industry
- Currently uses shared demo accounts (stale data, conflicts with other SEs)

**Goals**:
- ✅ Demo with realistic industry data (fintech, healthcare, retail)
- ✅ Show live data updates (add record → sync → appears in warehouse)
- ✅ Spin up new demo environment in < 60 seconds

**Frustrations**:
- 😡 "Shared demo account has wrong data (another SE modified it)"
- 😰 "Prospect asked 'Can you show healthcare data?' → No, only have generic data"
- ⏰ "Spent 2 hours setting up demo environment before customer call"

**What Sam Wants**:
> "Let me spin up a database with realistic healthcare data in 30 seconds. I need to demo in 10 minutes."

---

### Quaternary Persona: Jamie the DevRel Engineer

**Demographics**:
- Age: 29
- Role: Developer Relations Engineer at API company
- Salary: $130K base
- Content: Blog posts, video tutorials, conference talks

**Background**:
- Creates tutorials and demos for developers
- Needs realistic data that looks production-ready
- Currently uses Faker (but data is too generic to be compelling)

**Goals**:
- ✅ Create tutorials with realistic, relatable data
- ✅ Build demo apps that look production-ready
- ✅ Show best practices with real-world scenarios

**Frustrations**:
- 😡 "Tutorial says 'Employee 123' - looks unprofessional"
- 😰 "Conference demo with generic data → audience loses interest"
- ⏰ "Spent 4 hours creating realistic e-commerce data for blog post"

**What Jamie Wants**:
> "Give me realistic e-commerce data (products, orders, customers) for my tutorial. It needs to look professional."

---

## User Stories

### Epic 1: CLI Tool for Seeding Databases

**As a developer**, I want to seed my local database with realistic data in 30 seconds, so I can focus on building features instead of creating test data.

**User Stories**:
- [ ] As a backend dev, I want to run `sourcebox seed mysql --schema=fintech-loans --records=1000` and get realistic loan data
- [ ] As a data engineer, I want to seed Postgres with healthcare data for testing my ETL pipeline
- [ ] As a developer, I want to see progress (progress bar, ETA) while data is being seeded
- [ ] As a developer, I want to export SQL instead of inserting (for version control or sharing)
- [ ] As a developer, I want to list all available schemas before choosing one

**Acceptance Criteria**:
- ✅ CLI command completes in < 30 seconds (for 1,000 records)
- ✅ Data is realistic and verticalized (industry-specific)
- ✅ Supports MySQL + Postgres
- ✅ Progress bar shows status
- ✅ Can export to SQL file (`--output` flag)

---

### Epic 2: Docker Images for Quick Setup

**As a developer**, I want to spin up a pre-seeded database with Docker in 10 seconds, so I can test my app immediately.

**User Stories**:
- [ ] As a backend dev, I want to run `docker run sourcebox/mysql-fintech` and get a MySQL database with fintech data
- [ ] As a data engineer, I want to use Docker Compose to spin up multiple databases (fintech + healthcare)
- [ ] As a sales engineer, I want to spin up a demo environment for a customer in 30 seconds
- [ ] As a developer, I want to connect to the database immediately (no manual setup)

**Acceptance Criteria**:
- ✅ Docker container starts in < 10 seconds
- ✅ Database is pre-seeded with data (no manual seeding needed)
- ✅ 6 Docker images available (MySQL + Postgres, 3 industries each)
- ✅ Docker Compose file for multi-database setup
- ✅ Public on Docker Hub (no login required)

---

### Epic 3: Schema Library

**As a developer**, I want to browse available schemas and choose the one that fits my use case, so I can get started quickly.

**User Stories**:
- [ ] As a developer, I want to run `sourcebox list-schemas` and see all available schemas
- [ ] As a developer, I want to see schema details (tables, columns, record count) in docs
- [ ] As a developer, I want to see example queries for each schema (SQL snippets)
- [ ] As a developer, I want to know which industries are supported (fintech, healthcare, retail, etc.)

**Acceptance Criteria**:
- ✅ `sourcebox list-schemas` shows all schemas with descriptions
- ✅ Documentation site has schema library page
- ✅ Each schema shows: tables, columns, relationships, record count
- ✅ Example SQL queries provided for each schema

---

### Epic 4: Realistic, Verticalized Data

**As a developer**, I want data that looks like real industry data (not "John Doe"), so my demos are credible and relatable.

**User Stories**:
- [ ] As a fintech dev, I want loan amounts that follow realistic distributions ($5K-$500K, weighted toward $50K-$150K)
- [ ] As a healthcare dev, I want diagnoses that use real ICD-10 codes (diabetes, hypertension, etc.)
- [ ] As a retail dev, I want products that are real (iPhone 15 Pro, Nike Air Max) with realistic prices
- [ ] As a developer, I want relationships between tables (foreign keys, referential integrity)
- [ ] As a developer, I want edge cases (NULL values, extreme values) for testing

**Acceptance Criteria**:
- ✅ Data looks realistic (no "John Doe" or "Acme Corp")
- ✅ Industry-specific data (fintech = loans, healthcare = diagnoses)
- ✅ Realistic distributions (credit scores are bell-curved, not random)
- ✅ Foreign keys work (no orphaned records)
- ✅ Edge cases included (10% of data has unusual values for testing)

---

## Product Vision

### Vision Statement

> **SourceBox makes demo data effortless for developers, so they can focus on building great products instead of writing SQL.**

### 3-Year Vision

**Year 1: CLI + Docker MVP (Open Source)**
- Product: CLI tool + Docker images with 10 schemas (fintech, healthcare, retail, etc.)
- Distribution: npm, homebrew, Docker Hub
- Scale: 1,000 GitHub stars, 10,000 Docker Hub pulls
- Outcome: Standard tool for developers creating demo data

**Year 2: Cloud SaaS (Freemium)**
- Product: Cloud-hosted databases on-demand (API access, no local setup)
- Scale: 1,000 paying users, $20K MRR
- Outcome: "Docker for databases" - spin up realistic data in the cloud

**Year 3: Platform + Marketplace**
- Product: Community marketplace (contribute schemas), visual schema builder (drag-and-drop)
- Scale: 10,000 paying users, $100K MRR
- Outcome: Platform for verticalized demo data - acquired by Vercel/Supabase or IPO path

---

## MVP Scope

### What's IN Scope (MVP - 12 Weeks)

**Core Features**:
1. ✅ **CLI tool** (Go, distributed via npm + homebrew)
2. ✅ **3 verticalized schemas** (fintech-loans, healthcare-patients, retail-ecommerce)
3. ✅ **2 database types** (MySQL, Postgres)
4. ✅ **Docker images** (6 total: 3 schemas × 2 databases, published to Docker Hub)
5. ✅ **Documentation site** (Docusaurus, quick start guide, schema library)
6. ✅ **Open source** (MIT license, GitHub)

**Target Users**: Backend developers, data engineers (not SEs for MVP)

**Success Criteria**:
- 100+ GitHub stars in first week
- 500+ Docker Hub pulls in first month
- 20+ CLI installs in first week
- 5+ developers say "This saves me hours per week"

### What's OUT of Scope (Post-MVP)

**Phase 2 (Months 4-6)**:
- ❌ 7 more schemas (reach 10 total)
- ❌ MongoDB support
- ❌ Cloud SaaS (freemium)
- ❌ Custom schema builder

**Phase 3 (Year 2)**:
- ❌ Marketplace (community schemas)
- ❌ Visual schema builder (drag-and-drop)
- ❌ API mocks (Stripe, Twilio, GitHub)
- ❌ Enterprise features (SSO, air-gapped)

---

## Feature Requirements

### F1: CLI Tool (Core)

**Description**: Command-line tool to seed databases with verticalized data

**User Flow**:
1. Developer installs: `npm install -g sourcebox`
2. Developer runs: `sourcebox seed mysql --schema=fintech-loans --records=1000`
3. SourceBox connects to MySQL, creates tables, generates data, inserts records
4. Progress bar shows: "Seeding loans table... (500/1000) [███████___]"
5. Completion message: "✅ Seeded 1,000 records in 15 seconds"

**Technical Requirements**:
- **Language**: Go (single binary, cross-platform, fast)
- **CLI Framework**: Cobra (standard for Go CLIs)
- **Commands**:
  - `sourcebox seed` - Seed database with schema
  - `sourcebox list-schemas` - List available schemas
  - `sourcebox --version` - Show version
  - `sourcebox --help` - Show help
- **Flags for `seed`**:
  - `--database` (mysql, postgres) - Required
  - `--schema` (fintech-loans, etc.) - Required
  - `--records` (number, default 1000) - Optional
  - `--host`, `--port`, `--user`, `--password`, `--db-name` - Connection params
  - `--output` (export SQL instead of inserting) - Optional
  - `--dry-run` (show what would be seeded) - Optional
  - `--quiet` (suppress output except errors) - Optional
- **Distribution**:
  - npm package: `npm install -g sourcebox`
  - Homebrew: `brew install sourcebox`
  - Binary download: GitHub Releases (macOS, Linux, Windows)

**Acceptance Criteria**:
- ✅ CLI installs via npm + homebrew
- ✅ `seed` command completes in < 30 seconds (1,000 records)
- ✅ Progress bar shows status
- ✅ Supports MySQL + Postgres
- ✅ Error messages are clear and actionable
- ✅ `--help` shows all commands and flags

**Priority**: P0 (Must-Have for MVP)

---

### F2: Verticalized Schemas

**Description**: Industry-specific data schemas with realistic distributions

**Requirements**:

**Schema Format** (JSON):
```json
{
  "name": "fintech-loans",
  "description": "Loan records for fintech demos",
  "version": "1.0.0",
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {"name": "name", "type": "varchar(100)", "generator": "name"},
        {"name": "email", "type": "varchar(100)", "generator": "email", "unique": true},
        {"name": "credit_score", "type": "int", "generator": "credit_score"}
      ]
    },
    {
      "name": "loans",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {"name": "borrower_id", "type": "int", "foreign_key": "borrowers.id"},
        {"name": "amount", "type": "decimal(10,2)", "generator": "loan_amount"},
        {"name": "rate", "type": "decimal(5,2)", "generator": "interest_rate"},
        {"name": "status", "type": "varchar(20)", "generator": "loan_status"}
      ]
    }
  ]
}
```

**Data Generators** (Custom Logic):
- **Fintech**:
  - `loan_amount()`: $5K-$500K, weighted toward $50K-$150K (realistic distribution)
  - `interest_rate()`: 3%-15%, weighted toward 6%-10% (prime rates)
  - `credit_score()`: 300-850, bell curve around 680 (US average)
  - `loan_status()`: "active" (70%), "paid" (20%), "delinquent" (10%)

- **Healthcare**:
  - `diagnosis()`: Real ICD-10 codes (E11 = diabetes, I10 = hypertension)
  - `medication()`: Real drug names (Lipitor, Metformin, Lisinopril)
  - `visit_type()`: "routine checkup" (60%), "emergency" (20%), "follow-up" (20%)
  - `insurance_provider()`: Real insurers (Blue Cross, Aetna, UnitedHealthcare)

- **Retail**:
  - `product_name()`: Real products ("iPhone 15 Pro", "Nike Air Max")
  - `price()`: $5-$2,000, realistic for product category
  - `order_status()`: "pending" (10%), "shipped" (40%), "delivered" (40%), "returned" (10%)

**Acceptance Criteria**:
- ✅ 3 schemas available (fintech, healthcare, retail)
- ✅ Data looks realistic (no "John Doe" or "Acme Corp")
- ✅ Distributions are realistic (credit scores bell-curved, not random)
- ✅ Foreign keys work (referential integrity)
- ✅ Edge cases included (10% unusual values for testing)

**Priority**: P0 (Must-Have for MVP)

---

### F3: Docker Images

**Description**: Pre-seeded Docker images for instant database setup

**Requirements**:

**Image Structure**:
- **Base images**: `mysql:8.0`, `postgres:16`
- **Seed data**: Baked into image (no manual seeding)
- **Environment variables**:
  - `MYSQL_ROOT_PASSWORD=password`
  - `MYSQL_DATABASE=demo`
  - `POSTGRES_PASSWORD=password`
  - `POSTGRES_DB=demo`

**Images to Publish**:
1. `sourcebox/mysql-fintech:latest`
2. `sourcebox/mysql-healthcare:latest`
3. `sourcebox/mysql-retail:latest`
4. `sourcebox/postgres-fintech:latest`
5. `sourcebox/postgres-healthcare:latest`
6. `sourcebox/postgres-retail:latest`

**Docker Compose Example**:
```yaml
version: '3.8'
services:
  mysql-fintech:
    image: sourcebox/mysql-fintech:latest
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: demo

  postgres-healthcare:
    image: sourcebox/postgres-healthcare:latest
    ports:
      - "5432:5432"
    environment:
      POSTGRES_PASSWORD: password
      POSTGRES_DB: demo
```

**Technical Requirements**:
- **Build**: Multi-stage builds (optimize size)
- **Size**: < 500MB per image (compressed)
- **Startup**: < 10 seconds (container ready)
- **Distribution**: Docker Hub (public, no login required)

**Acceptance Criteria**:
- ✅ 6 images published to Docker Hub
- ✅ Images are public (no authentication needed)
- ✅ Container starts in < 10 seconds
- ✅ Database is pre-seeded (ready to query)
- ✅ Docker Compose file works (multi-database setup)

**Priority**: P0 (Must-Have for MVP)

---

### F4: Documentation Site

**Description**: Comprehensive documentation for developers (getting started, schemas, CLI reference)

**Requirements**:

**Tech Stack**: Docusaurus (standard for dev docs)

**Pages**:
1. **Home**: Hero, quick start (5-minute guide), use cases
2. **Getting Started**: Install CLI, seed database, query data
3. **CLI Reference**: All commands, flags, examples
4. **Schema Library**: All schemas, table definitions, example queries
5. **Docker Images**: All images, usage examples, Docker Compose
6. **FAQ**: Common questions (How to customize? How to contribute?)
7. **Contributing**: How to add new schemas, report bugs

**Quick Start Example** (Home Page):
```bash
# Install
npm install -g sourcebox

# Seed MySQL with fintech data
sourcebox seed mysql --schema=fintech-loans --records=1000

# Connect and query
mysql -u root -p demo
mysql> SELECT * FROM loans LIMIT 5;
```

**Schema Library Example** (Per Schema):
- **Name**: Fintech Loans
- **Description**: Loan records for fintech demos
- **Tables**: loans, borrowers, payments, credit_scores
- **Record Count**: 1,000 (default)
- **Use Cases**: Lending apps, banking demos, payment systems
- **Example Queries**:
  ```sql
  -- Top 10 borrowers by credit score
  SELECT name, credit_score FROM borrowers ORDER BY credit_score DESC LIMIT 10;

  -- Total loan amount by status
  SELECT status, SUM(amount) FROM loans GROUP BY status;
  ```

**Acceptance Criteria**:
- ✅ Documentation site live (Docusaurus on Vercel)
- ✅ Quick start guide (5-minute guide to first seed)
- ✅ CLI reference (all commands documented)
- ✅ Schema library (all 3 schemas documented)
- ✅ Docker images page (all 6 images documented)
- ✅ FAQ page (common questions answered)

**Priority**: P0 (Must-Have for MVP)

---

### F5: Open Source Repository

**Description**: GitHub repository with code, examples, issues, contributions

**Requirements**:

**Repository Structure**:
```
/sourcebox
  /cmd          # CLI commands (Cobra)
  /pkg          # Core logic (generators, schemas, database)
  /schemas      # JSON schema files
  /docker       # Dockerfiles for each image
  /docs         # Docusaurus site
  /examples     # Example projects (Next.js, Python, etc.)
  README.md     # Quick start, badges, use cases
  LICENSE       # MIT license
  CONTRIBUTING.md
  CODE_OF_CONDUCT.md
```

**README.md** (Key Sections):
- Hero: "Realistic, verticalized demo data for developers"
- Badges: Stars, Docker pulls, npm downloads, license
- Quick Start: Install, seed, query (3 steps, 30 seconds)
- Use Cases: Backend dev, data engineer, SE, DevRel, QA
- Available Schemas: Table (10 schemas, 3 available now)
- Docker Images: Table (6 images)
- Contributing: How to contribute schemas, report bugs
- Legal Note: "Developed independently on personal equipment"

**Issues & Discussions**:
- Issue templates: Bug report, feature request, schema request
- Discussions: Q&A, show and tell, ideas

**Acceptance Criteria**:
- ✅ GitHub repository public (MIT license)
- ✅ README.md comprehensive (quick start, use cases, schemas)
- ✅ CONTRIBUTING.md (how to contribute)
- ✅ LICENSE (MIT)
- ✅ Issue templates (bug, feature, schema request)
- ✅ Legal note (developed independently)

**Priority**: P0 (Must-Have for MVP)

---

## Non-Functional Requirements

### Performance

| Metric | Target | Critical Path |
|--------|--------|---------------|
| **CLI seed time** | < 30s (1,000 records) | Fast enough to not disrupt workflow |
| **Docker startup time** | < 10s | Instant feel |
| **Data generation rate** | 100+ records/sec | Batch inserts, parallel processing |
| **Binary size** | < 50MB | Small download, fast install |

### Scalability

| Metric | Target | Rationale |
|--------|--------|-----------|
| **Records per schema** | 1,000-10,000 | Balance between realism and speed |
| **Schemas supported** | 10+ (MVP: 3) | Room for growth |
| **Concurrent users** | 1,000+ | Open source, anyone can use |

### Reliability

| Metric | Target | Impact if Missed |
|--------|--------|------------------|
| **CLI success rate** | 99%+ | Broken CLI → bad first impression |
| **Docker image availability** | 99.9% (Docker Hub SLA) | Downtime → can't pull images |
| **Data quality** | 95%+ realistic | Generic data → poor demos |

### Security

| Requirement | Implementation | Priority |
|-------------|----------------|----------|
| **No PII** | Only synthetic data, no real names/emails | P0 (must-have) |
| **No secrets** | No API keys, passwords in code | P0 (must-have) |
| **Safe defaults** | Default passwords documented, easy to change | P1 (post-MVP) |
| **Audit trail** | (Not needed for MVP - no user accounts) | P2 (Phase 2) |

### Cost

| Resource | Target Cost | Budget/Month |
|----------|-------------|--------------|
| **Development** | $0 (open source) | $0 |
| **Hosting** (docs site) | $0 (Vercel free tier) | $0 |
| **Docker Hub** | $0 (free public repos) | $0 |
| **Total Infrastructure** | | **$0/mo** (MVP) |

**Unit Economics (Phase 2 Cloud SaaS)**:
- $20/mo subscription - $2/mo cost (cloud DB) = $18/mo gross margin (90%)

---

## User Experience

### Design Principles

1. **Developer-First**: CLI > GUI, terminal-friendly, works with existing tools
2. **Zero Config**: Works out of the box, no complex setup
3. **Fast**: 30 seconds from install to seeded database
4. **Boring Tech**: Proven tools (Go, Docker, MySQL/Postgres), no bleeding edge
5. **Transparent**: Clear progress, helpful errors

### Key User Flows

#### Flow 1: Install and Seed (First Time User)

**Scenario**: Alex is a backend dev building a fintech app, needs realistic loan data

**Steps**:
1. Alex searches Google: "realistic fintech demo data"
2. Finds SourceBox GitHub repo → Clicks README
3. Sees Quick Start:
   ```bash
   npm install -g sourcebox
   sourcebox seed mysql --schema=fintech-loans --records=1000
   ```
4. Copies command, runs in terminal
5. Sees progress bar: "Seeding loans table... (500/1000) [███████___] ETA 15s"
6. Completion: "✅ Seeded 1,000 records in 15 seconds"
7. Connects to MySQL: `mysql -u root -p demo`
8. Queries: `SELECT * FROM loans LIMIT 5;` → Sees realistic fintech loans
9. Alex: "This is exactly what I needed!"

**Success Metric**: 90%+ of users complete this flow without support

---

#### Flow 2: Docker Quick Start

**Scenario**: Maria is a data engineer testing an ETL pipeline, needs healthcare data

**Steps**:
1. Maria needs healthcare data for testing
2. Searches: "healthcare demo data docker"
3. Finds SourceBox docs → Docker Images page
4. Copies command:
   ```bash
   docker run -p 5432:5432 sourcebox/postgres-healthcare:latest
   ```
5. Runs in terminal → Container starts in 10 seconds
6. Connects to Postgres: `psql -h localhost -U postgres -d demo`
7. Queries: `SELECT * FROM patients LIMIT 5;` → Sees realistic healthcare data
8. Maria: "Perfect for my pipeline test!"

**Success Metric**: 95%+ of Docker pulls result in successful container startup

---

#### Flow 3: Exploring Schema Library

**Scenario**: Sam is a sales engineer preparing a demo, needs to choose right schema

**Steps**:
1. Sam needs realistic data for customer demo (retail industry)
2. Runs: `sourcebox list-schemas`
3. Sees output:
   ```
   SCHEMA                  TABLES                         RECORDS
   fintech-loans          loans, borrowers, payments      1,000
   healthcare-patients    patients, visits, claims        500
   retail-ecommerce       products, orders, customers     2,000
   ```
4. Picks retail-ecommerce → Visits docs for details
5. Docs show: tables, columns, example queries
6. Sam copies example query for demo
7. Runs: `sourcebox seed postgres --schema=retail-ecommerce --records=2000`
8. Demo ready in 30 seconds

**Success Metric**: 80%+ of users visit docs before first seed

---

### CLI Output Examples

**Success Output**:
```bash
$ sourcebox seed mysql --schema=fintech-loans --records=1000

Connecting to MySQL (localhost:3306)...
✓ Connected

Creating tables...
✓ Created 4 tables (borrowers, loans, payments, credit_scores)

Seeding data...
⣾ Seeding borrowers table... (250/250) [██████████] 100% | ETA: 0s
⣾ Seeding loans table... (1000/1000) [██████████] 100% | ETA: 0s
⣾ Seeding payments table... (3500/3500) [██████████] 100% | ETA: 0s
⣾ Seeding credit_scores table... (250/250) [██████████] 100% | ETA: 0s

✅ Successfully seeded 4,950 records in 15 seconds

Tables:
  - borrowers: 250 records
  - loans: 1,000 records
  - payments: 3,500 records
  - credit_scores: 250 records

Connect: mysql -u root -p demo
```

**Error Output** (Clear, Actionable):
```bash
$ sourcebox seed mysql --schema=fintech-loans --records=1000

Connecting to MySQL (localhost:3306)...
✗ Connection failed: Can't connect to MySQL server on 'localhost:3306'

Possible causes:
  1. MySQL is not running (start with: mysql.server start)
  2. Wrong host or port (use --host and --port flags)
  3. Wrong credentials (use --user and --password flags)

For help, run: sourcebox --help
```

---

## Success Metrics

### North Star Metric

**Hours saved per developer per week** (Target: 5+ hours)

**Measurement**: Community surveys, GitHub issues, testimonials

### Key Performance Indicators (KPIs)

#### Product Metrics

| Metric | Target (Week 1) | Target (Month 1) | Target (Month 3) |
|--------|-----------------|------------------|------------------|
| **GitHub stars** | 100+ | 500+ | 1,000+ |
| **Docker Hub pulls** | 500+ | 5,000+ | 20,000+ |
| **CLI installs** (npm + homebrew) | 20+ | 200+ | 1,000+ |
| **Documentation visits** | 1,000+ | 10,000+ | 50,000+ |

#### Engagement Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| **CLI usage** | 50+ developers actively using | GitHub telemetry (opt-in) |
| **Docker usage** | 500+ container starts/week | Docker Hub stats |
| **Community contributions** | 5+ schema requests | GitHub issues |
| **NPS** | 40+ (good for MVP) | Community survey |

#### Business Metrics (Phase 2 - Cloud SaaS)

| Metric | Target (Month 6) | Target (Month 12) |
|--------|------------------|-------------------|
| **Paying users** | 100 | 1,000 |
| **MRR** | $2,000 | $20,000 |
| **Free → Paid conversion** | 5% | 10% |
| **Churn** | < 10%/month | < 5%/month |

---

## Competitive Analysis

### Current Landscape

**Key Insight**: There is NO direct competitor for verticalized, realistic demo data. The market is fragmented:

| Solution | What They Do | Limitations |
|----------|-------------|-------------|
| **Faker.js** | Generate random data (names, emails, etc.) | Too generic ("John Doe", "Acme Corp"), no verticalization |
| **Mockaroo** | Web-based data generator | Web UI (not CLI/Docker), limited schema customization, not verticalized |
| **Manual SQL** | Write INSERT statements | Tedious (1,000+ lines), time-consuming (hours), error-prone |
| **Production dumps** | Export real data | Security risk (PII), GDPR/HIPAA nightmare, data drift |

### Competitive Matrix

| Feature | SourceBox | Faker.js | Mockaroo | Manual SQL | Production Dumps |
|---------|-----------|----------|----------|------------|------------------|
| **Verticalized** | ✅ Yes (industry-specific) | ❌ No (generic) | ⚠️ Partial | ✅ Yes (custom) | ✅ Yes (real data) |
| **Speed** | ✅ 30 seconds | ✅ Instant | ⚠️ Minutes (web UI) | ❌ Hours | ⚠️ Minutes (export + import) |
| **Ease of Use** | ✅ CLI + Docker | ✅ API/library | ⚠️ Web UI (no CI/CD) | ❌ Manual coding | ⚠️ Manual export |
| **Realistic Data** | ✅ Yes | ❌ No | ⚠️ Partial | ✅ Yes (if you code it) | ✅ Yes (real data) |
| **Security** | ✅ Synthetic (no PII) | ✅ Synthetic | ✅ Synthetic | ✅ Synthetic | ❌ Risk (PII) |
| **Cost** | ✅ Free (open source) | ✅ Free | ⚠️ $60/mo (paid plans) | ✅ Free | ⚠️ Risk (compliance cost) |

### Unique Value Proposition

**SourceBox is the ONLY tool that combines:**
1. **Verticalized data** (fintech looks like fintech, healthcare looks like healthcare)
2. **Developer-first** (CLI + Docker, works in CI/CD)
3. **Fast** (30 seconds, not hours)
4. **Open source** (free forever, MIT license)

**Why This Matters**:
- **Verticalized** → Demos resonate with prospects (not "John Doe")
- **Developer-first** → Integrates with existing workflows (not web UIs)
- **Fast** → No time wasted (30 seconds vs hours of manual SQL)
- **Open source** → Anyone can use, no lock-in

**Key Differentiators**:
1. **Solves BOTH problems** (realism + speed) - no one else does
2. **Built for developers** (CLI + Docker, not web UIs)
3. **Open source** (MIT license, community-driven)
4. **Vertical-first** (industry-specific, not generic)

---

## Go-To-Market

### Target Market (ICP)

**Phase 1 (MVP - Open Source)**:
- **Target**: All developers (backend, data engineers, SEs, DevRel, QA)
- **TAM**: 1M+ developers globally who need demo data
- **Distribution**: Open source (GitHub, Docker Hub, npm)
- **Revenue**: $0 (free forever for CLI + Docker)

**Phase 2 (Cloud SaaS - Months 7-12)**:
- **Target**: Teams (10+ developers) who want cloud-hosted databases
- **ICP**: SaaS companies, agencies, data platform companies
- **Pain**: Local setup is slow, want on-demand databases
- **Willingness to Pay**: $20-200/mo (freemium)

### GTM Strategy

#### Phase 1: Open Source Launch (Months 1-3)

**Goal**: 100+ GitHub stars, 500+ Docker Hub pulls in first week

**Approach**:
1. **Pre-launch** (Week 11):
   - Alpha test with 10 developers
   - Collect testimonials ("SourceBox saved me 8 hours this week")
   - Prepare launch posts (HN, Reddit, LinkedIn, dev.to)

2. **Launch Day** (Week 12, Friday):
   - 9 AM ET: Hacker News "Show HN" post
   - 12 PM ET: Reddit r/programming post
   - 3 PM ET: LinkedIn personal post
   - 6 PM ET: dev.to article
   - Monitor all day, reply to all comments

3. **Post-launch** (Week 13+):
   - ProductHunt launch (Week 13, Tuesday)
   - Cross-post to r/devops, r/webdev
   - Twitter threads (2-3 per week)
   - Answer questions on Stack Overflow

**Launch Post Example** (Hacker News):
```
Show HN: SourceBox - Realistic demo data for developers (CLI + Docker)

I'm a solutions architect, and I was frustrated with how long it takes to
create realistic demo data. Faker gives you "John Doe", and manual SQL is tedious.

So I built SourceBox - a CLI tool and Docker images with verticalized, realistic
data for fintech, healthcare, retail, and more.

Examples:
- CLI: `sourcebox seed mysql --schema=fintech-loans --records=1000` (30 seconds)
- Docker: `docker run sourcebox/mysql-fintech` (10 seconds)

Open source (MIT), free forever. Cloud version (freemium) coming in Phase 2.

Would love your feedback! GitHub: [link]
```

**Success Metric**: 100+ GitHub stars, 500+ Docker Hub pulls, 50+ comments/feedback

---

#### Phase 2: Product-Led Growth (Months 4-12)

**Goal**: 1,000 paying users, $20K MRR by Month 12

**Approach**:
1. **Free tier** (CLI + Docker - forever free):
   - Unlimited schemas, unlimited records
   - Open source, self-hosted
   - No lock-in, no credit card

2. **Cloud tier** ($20/mo per user):
   - Cloud-hosted databases (on-demand)
   - API access (CI/CD integration)
   - 7-day retention
   - 5 databases active at once

3. **Team tier** ($200/mo for 10 users):
   - 30-day retention
   - 50 databases active at once
   - Team templates (shared schemas)
   - Priority support

4. **Viral loop**:
   - "Share this schema" button → generate URL → new signups
   - "Built with SourceBox" badge in README → backlink

**Pricing Page** (Phase 2):
```
FREE                    CLOUD ($20/mo)          TEAM ($200/mo)
----                    ----                    ----
✓ CLI + Docker          ✓ Cloud-hosted DBs      ✓ Everything in Cloud
✓ All schemas           ✓ API access            ✓ 30-day retention
✓ Unlimited records     ✓ 7-day retention       ✓ 50 databases
✓ Open source           ✓ 5 databases           ✓ Team templates
                        ✓ CI/CD integration     ✓ Priority support
```

**Success Metric**: 1,000 paying users, $20K MRR by Month 12

---

#### Phase 3: Community & Content (Ongoing)

**Goal**: Build brand awareness in developer community

**Channels**:
1. **Hacker News**: Weekly "Show HN" updates (new schemas, features)
2. **Reddit**: r/programming, r/devops, r/webdev (case studies, AMAs)
3. **dev.to**: Weekly articles ("How to build a fintech app in 1 hour")
4. **Twitter**: Daily tips, schema announcements, user testimonials
5. **ProductHunt**: Launch updates (new schemas, cloud version)

**Content**:
- Blog: "Developer's Guide to Realistic Demo Data" (SEO, lead magnet)
- Video: "Build a SaaS app with realistic data in 10 minutes"
- Podcast: Guest on developer podcasts (Syntax.fm, Software Engineering Daily)

**Success Metric**: 40%+ signups from community/content by Month 12

---

## Roadmap

### MVP (Months 1-3): Build & Launch

**Week 1-2: Validation**
- [ ] Interview 20 developers (backend, data engineers, SEs)
- [ ] Validate problem (how much time wasted? would they use CLI/Docker?)
- [ ] Pre-sell to 10 alpha testers (commit to feedback)

**Week 3-8: Build Core**
- [ ] CLI tool (Go + Cobra)
- [ ] Data generation engine (Faker + custom generators)
- [ ] 3 schemas (fintech, healthcare, retail)
- [ ] Docker images (6 total: MySQL + Postgres × 3)
- [ ] Documentation site (Docusaurus)

**Week 9-10: Alpha Testing**
- [ ] Deploy to 10 alpha testers
- [ ] Measure: time saved, data quality, ease of use
- [ ] Iterate based on feedback

**Week 11-12: Polish & Launch**
- [ ] Fix top 10 bugs from alpha
- [ ] Publish to npm + homebrew + Docker Hub
- [ ] Open source launch (GitHub, HN, Reddit)

**Success Criteria**:
- ✅ 100+ GitHub stars in first week
- ✅ 500+ Docker Hub pulls in first month
- ✅ 20+ CLI installs in first week
- ✅ 5+ developers say "This saves me hours per week"

---

### Phase 2 (Months 4-6): Scale Open Source

**Month 4: Add More Schemas**
- [ ] Add 7 more schemas (reach 10 total):
  - SaaS (users, subscriptions, usage)
  - E-commerce (orders, products, reviews)
  - Logistics (shipments, routes, drivers)
  - Education (students, courses, grades)
  - Real Estate (properties, agents, transactions)
  - Insurance (policies, claims, premiums)
  - Manufacturing (products, orders, suppliers)

**Month 5: MongoDB Support**
- [ ] Add MongoDB support (CLI + Docker)
- [ ] Create 3 MongoDB schemas (e-commerce, SaaS, logs)
- [ ] Update docs with MongoDB examples

**Month 6: Community Growth**
- [ ] Schema request workflow (community can request)
- [ ] Contributor guide (how to add new schemas)
- [ ] Community showcase (projects built with SourceBox)

**Success Criteria**:
- ✅ 1,000+ GitHub stars
- ✅ 10,000+ Docker Hub pulls
- ✅ 10+ community schema requests

---

### Phase 3 (Months 7-12): Cloud SaaS

**Month 7-9: Build Cloud MVP**
- [ ] API endpoint: `POST /api/databases` (create cloud DB)
- [ ] Seed database with schema
- [ ] Return connection string
- [ ] Auto-delete after TTL (7 days)

**Month 10-12: Launch Freemium**
- [ ] Pricing page ($0 → $20/mo → $200/mo)
- [ ] Stripe integration (payment, subscriptions)
- [ ] Self-serve signup (no sales calls)
- [ ] Email onboarding sequence

**Success Criteria**:
- ✅ 100+ paying users
- ✅ $2,000 MRR
- ✅ 5%+ free → paid conversion
- ✅ < 10% monthly churn

---

### Phase 4 (Year 2): Platform + Marketplace

**Q1-Q2: Marketplace**
- [ ] Community schema marketplace
- [ ] Submit your own schemas
- [ ] Upvote/downvote schemas
- [ ] Featured schemas (curated)

**Q3-Q4: Visual Schema Builder**
- [ ] Drag-and-drop schema builder
- [ ] Define tables, columns, relationships
- [ ] Generate schema JSON
- [ ] Export to CLI or publish to marketplace

**Success Criteria**:
- ✅ 1,000+ paying users
- ✅ $20,000 MRR
- ✅ 100+ community schemas

---

## Dependencies & Risks

### Critical Dependencies

| Dependency | Risk | Mitigation |
|------------|------|------------|
| **Go ecosystem** | If CLI framework breaks | Use mature libraries (Cobra, go-faker) |
| **Docker Hub** | If downtime or limits hit | Multi-registry support (GitHub Container Registry) |
| **npm/Homebrew** | If package rejected or delayed | Direct binary download (GitHub Releases) |
| **Community adoption** | If no one uses it | Validate problem first (20 interviews), pre-sell alpha testers |

### Key Risks

**Risk 1: Problem Not Deep Enough**
- **Description**: Developers don't care enough about demo data to switch tools
- **Impact**: High (no adoption)
- **Probability**: Low (validated in interviews)
- **Mitigation**: Interview 20 developers first, pre-sell to 10 alpha testers

**Risk 2: Technical Feasibility**
- **Description**: Can't achieve 30-second seed time (too slow)
- **Impact**: Medium (poor UX)
- **Probability**: Low (Go is fast, batch inserts)
- **Mitigation**: Benchmark early, optimize if needed (accept 60s for MVP)

**Risk 3: Competitive Response**
- **Description**: Faker adds verticalization, or new competitor launches
- **Impact**: Medium (lose differentiation)
- **Probability**: Medium (6-12 months lead time)
- **Mitigation**: Move fast (ship MVP in 12 weeks), build moat (community, schemas)

**Risk 4: Legal/IP**
- **Description**: Current employer claims IP or non-compete violation
- **Impact**: High (can't launch)
- **Probability**: Low (independent development, no employer resources)
- **Mitigation**: Build outside work hours, no employer resources, consult lawyer, add legal note to docs

**Risk 5: Market Size Too Small**
- **Description**: Only 10K developers need demo data, not 1M+
- **Impact**: High (can't scale)
- **Probability**: Low (validated TAM = 1M+ devs)
- **Mitigation**: Start with open source (free), expand to cloud (freemium), then platform (marketplace)

---

## Appendix

### Glossary

- **CLI**: Command-Line Interface (terminal tool)
- **Schema**: Database structure (tables, columns, relationships)
- **Verticalized**: Industry-specific (fintech, healthcare, retail)
- **Synthetic Data**: Fake data that looks real (no PII)
- **PII**: Personally Identifiable Information (names, emails, SSNs)
- **Faker**: Library for generating random data (generic)
- **Docker**: Containerization platform (package apps + dependencies)
- **Open Source**: Free software with public code (MIT license)

### Research Sources

1. **Primary Research**: 20+ developer interviews (backend devs, data engineers, SEs)
2. **Secondary Research**:
   - Stack Overflow Developer Survey (2024)
   - GitHub State of the Octoverse (2024)
   - dev.to community discussions (2024-2025)
   - Reddit r/programming pain point analysis (2024-2025)

### Comparable Companies (Inspiration, Not Competitors)

- **Faker.js**: 40K+ GitHub stars, standard for random data generation
- **Mockaroo**: 500K+ users, web-based data generator
- **Docker**: 1B+ pulls/month, standard for containerization
- **Prisma**: 30K+ GitHub stars, developer-first database tool

---

## Legal Note

**This project is developed independently** on personal equipment, outside of work hours, with no use of employer resources or proprietary information.

All references to specific companies or products in documentation are for illustrative purposes only. SourceBox is a standalone, open source project with no affiliation to any employer or third party.

---

**Document Version**: 2.0 (Developer Platform Edition)
**Last Updated**: 2025-01-14
**Next Review**: 2025-02-01 (after MVP validation)
**Owner**: Founder / Product Lead
