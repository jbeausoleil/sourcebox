# SourceBox: Project Brief

> **One-liner**: Realistic, verticalized demo data for developers - spin up production-like databases in 30 seconds.

## 🎯 The Problem

**Developers waste 5-10 hours per week creating realistic demo data**

### The Current State of Demo Data

**Option 1: Generic Faker libraries**
```javascript
// What you write
const user = faker.person.fullName()  // "John Doe"
const company = faker.company.name()  // "Acme Corp"

// Problem: NOT verticalized
// - Healthcare dev needs patient records, not "Acme Corp"
// - Fintech dev needs loan applications, not generic users
// - Retail dev needs e-commerce orders, not random data
```

**Option 2: Manual SQL inserts**
```sql
-- Spend 3 hours writing this
INSERT INTO loans (amount, rate, status) VALUES
  (50000, 5.5, 'approved'),
  (25000, 6.2, 'pending'),
  -- ... 998 more rows

-- Problem: Tedious, not realistic, doesn't scale
```

**Option 3: Production data dumps**
```bash
# Download prod snapshot, scrub PII
# Problem: Security risk, compliance nightmare, still not perfect for demos
```

### Who This Affects

| Role | Use Case | Time Wasted |
|------|----------|-------------|
| **Backend developers** | Local development, testing | 2-3 hrs/week |
| **Data engineers** | Pipeline testing, demos | 5-10 hrs/week |
| **Sales Engineers** | Product demos, POCs | 10-15 hrs/week |
| **QA engineers** | Test data setup | 3-5 hrs/week |
| **DevRel / Solutions** | Demo apps, tutorials | 5-8 hrs/week |

**Total pain:**
- 🕐 **5-10 hours per developer per week** creating demo data
- 💸 **$100-200/week in lost productivity** per developer
- 😡 **Frustration** using unrealistic data ("Employee 123", "test@test.com")
- 🐛 **Bugs missed** because test data doesn't match real-world patterns

### The Core Insight

**There's no `docker run mysql-with-fintech-data` equivalent**

```bash
# This EXISTS:
docker run mysql  # Empty database in 30 seconds ✅

# This DOESN'T EXIST:
docker run mysql-with-fintech-loans-data  # 1,000 realistic loan records ❌
docker run postgres-with-healthcare-patients  # 500 patient records ❌
docker run mongodb-with-ecommerce-orders  # 10,000 order records ❌
```

**SourceBox fills this gap.**

---

## 💡 The Solution

**SourceBox = CLI tool + Docker images + verticalized schemas**

### Developer Experience

```bash
# Install once
npm install -g sourcebox
# or
brew install sourcebox

# Seed MySQL with realistic fintech data
sourcebox seed mysql --schema=fintech-loans --records=1000

# Output:
✅ Seeded MySQL database 'demo' with 1,000 loan records
✅ Tables: loans, borrowers, payments, credit_scores
✅ Data: Realistic names, addresses, credit scores, loan amounts
✅ Time: 15 seconds

# Or use Docker Compose
docker-compose up -d  # Uses sourcebox/mysql-fintech:latest
```

### What Gets Created

**Fintech loans schema example:**

| Table | Records | Realism |
|-------|---------|---------|
| `borrowers` | 1,000 | Real names, addresses (Faker), credit scores (distribution curve) |
| `loans` | 1,000 | Amounts ($5K-$500K), rates (3%-15%), status (approved/denied) |
| `payments` | 5,000 | Payment history with realistic dates, late payments |
| `credit_scores` | 1,000 | FICO scores (300-850) with realistic distribution |

**Healthcare patients schema example:**

| Table | Records | Realism |
|-------|---------|---------|
| `patients` | 500 | Real names, DOB, insurance info |
| `visits` | 2,000 | Visit dates, diagnoses (ICD-10 codes), treatments |
| `prescriptions` | 3,000 | Medication names, dosages, dates |
| `insurance_claims` | 1,500 | Claim amounts, statuses, dates |

### Core Value Props

1. **30-second setup** - `sourcebox seed` or `docker run`
2. **Verticalized schemas** - 10+ industries (fintech, healthcare, retail, SaaS, e-commerce, etc.)
3. **Realistic data** - Proper distributions, relationships, edge cases
4. **Multiple databases** - MySQL, Postgres, MongoDB, SQLite
5. **Local-first** - Runs on your laptop, no cloud required
6. **Open source** - MIT license, community-driven

---

## 📊 Market Opportunity

### ICP (Ideal Customer Profile)

**Primary: Backend / Data Engineers at B2B SaaS**
- **Company size**: Series A-C (50-500 employees)
- **Use case**: Demo data for local dev, testing, demos
- **Pain**: Waste hours creating fake data
- **Willingness to pay**: $20/mo for cloud version (vs $100+/week in time wasted)

**Secondary: Sales Engineers / Solutions Architects**
- **Company size**: Any B2B SaaS with technical sales
- **Use case**: Product demos, POCs, training
- **Pain**: Need realistic data that matches prospect's industry
- **Willingness to pay**: $50-200/mo per SE

### TAM (Total Addressable Market)

#### Wave 1: Developers at Data-Adjacent Companies (Years 1-2)

| Segment | Companies | Developers | Market Size |
|---------|-----------|------------|-------------|
| **Data integration** (Airbyte, Hightouch, etc.) | 20 | 2,000 | $2M ARR |
| **BI tools** (Tableau, Looker, Mode, Hex) | 30 | 5,000 | $5M ARR |
| **Databases** (Snowflake, Databricks, SingleStore) | 20 | 3,000 | $3M ARR |
| **ETL/Transform** (dbt Labs, Matillion) | 15 | 2,000 | $2M ARR |
| **API platforms** (MuleSoft, Zapier, Tray.io) | 25 | 5,000 | $5M ARR |
| **Total** | ~110 | ~17,000 | **$17M ARR** |

*At $10/developer/month (cloud version), ~50% adoption*

#### Wave 2: Backend Developers (Years 2-3)

| Segment | Developers | Adoption | Market Size |
|---------|------------|----------|-------------|
| **Backend devs at B2B SaaS** | 500,000 | 5% | $30M ARR |
| **Data engineers** | 100,000 | 10% | $12M ARR |
| **QA engineers** | 200,000 | 3% | $7M ARR |
| **Total** | 800,000 | ~6% avg | **$49M ARR** |

#### Wave 3: All Developers (Years 3-5)

| Segment | Market | Opportunity |
|---------|--------|-------------|
| **Open source (free tier)** | 10M+ developers | $0 (growth driver) |
| **Cloud version (paid)** | 1M developers @ $20/mo | $240M ARR |
| **Enterprise (team licenses)** | 10K companies @ $5K/year | $50M ARR |
| **Total TAM** | - | **$290M+ ARR** |

### Revenue Model

**Freemium Model:**

| Tier | Price | Features | Target Users |
|------|-------|----------|--------------|
| **Free** | $0 | CLI tool (open source), Docker images, 5 schemas | Individual developers |
| **Cloud** | $20/mo | Cloud-hosted instances, 50 schemas, API access, CI/CD integrations | Teams (5-20 devs) |
| **Team** | $200/mo | Unlimited schemas, custom schemas, priority support | Companies (20-100 devs) |
| **Enterprise** | Custom | Private deployment, air-gapped, custom integrations, SLA | Large enterprises (100+ devs) |

**Unit Economics:**

| Metric | Value | Notes |
|--------|-------|-------|
| **Free → Cloud conversion** | 5% | Industry standard for dev tools |
| **Cloud → Team conversion** | 20% | As teams grow |
| **ARPU (blended)** | $15/mo | Weighted average across tiers |
| **CAC** | $50 | Open source → organic → low CAC |
| **LTV** | $900 | 5 years × $15/mo × 12 months |
| **LTV:CAC** | 18:1 | Excellent for SaaS |

### Funding Requirements

**Bootstrap Path (Recommended):**
- **Investment**: $0 (develop nights/weekends)
- **Timeline**: 3 months to MVP, 6 months to first $1K MRR
- **Outcome**: Profitable by Month 12

**Venture Path (Alternative):**
- **Amount**: $500K pre-seed
- **Use**: Quit job, hire 1 engineer, 12 months runway
- **Outcome**: $50K MRR by Month 18, raise seed round

---

## 🏆 Why This Wins

### Unique Advantages

✅ **Unsolved problem** - No one does verticalized demo data as a service
✅ **Developer-first** - CLI tool, open source, Docker-native
✅ **Network effects** - More schemas → more value → more users → more schemas
✅ **PLG motion** - Free tier drives adoption, paid for cloud/teams
✅ **Massive TAM** - Millions of developers need this
✅ **Low CAC** - Open source → organic growth
✅ **Strong retention** - Embedded in dev workflow

### Competitive Moat (Years 1-3)

**Year 1: Schema Library**
- 50+ verticalized schemas (fintech, healthcare, retail, SaaS, etc.)
- Would take competitors 12-18 months to replicate

**Year 2: Community**
- 10K+ GitHub stars
- Community-contributed schemas (marketplace)
- Network effects kick in

**Year 3: Platform**
- Integrations (CI/CD, Kubernetes, Terraform)
- Custom schema builder (GUI)
- Enterprise features (private schemas, air-gapped)

---

## 🎨 Product Vision

### Phase 1: CLI + Docker (Months 1-6) - Open Source MVP

**Problems solved:**
1. Developers waste hours creating demo data
2. Faker is too generic, SQL inserts are tedious
3. No `docker run` equivalent for realistic data

**Solution:** CLI tool + Docker images with verticalized schemas

**Features:**
- ✅ CLI tool (`sourcebox seed <db> --schema=<industry>`)
- ✅ Docker images (pre-seeded databases on Docker Hub)
- ✅ 10 schemas: fintech, healthcare, retail, SaaS, e-commerce, logistics, education, real estate, insurance, manufacturing
- ✅ 3 databases: MySQL, Postgres, MongoDB
- ✅ Open source (MIT license)
- ✅ Documentation site
- ✅ Docker Compose examples

**Success metrics:**
- 1,000 GitHub stars
- 10,000 Docker Hub pulls
- 100 daily CLI installs

---

### Phase 2: Cloud SaaS (Months 7-12) - Freemium Launch

**Problem:** Developers want cloud-hosted instances for CI/CD, testing

**Solution:** Cloud-hosted SourceBox with API

**Features:**
- ✅ Cloud-hosted databases (on-demand spin-up)
- ✅ API access (`POST /databases` → get connection string)
- ✅ CI/CD integrations (GitHub Actions, GitLab CI)
- ✅ 50+ schemas (expand library)
- ✅ Custom schema builder (JSON/YAML templates)
- ✅ Team collaboration (shared schemas)
- ✅ Pricing tiers (Free, Cloud, Team)

**Success metrics:**
- $10K MRR
- 1,000 paying users
- 5% free → paid conversion

---

### Phase 3: Platform + Marketplace (Year 2) - Scale to $100K MRR

**Problem:** Every company has unique data needs

**Solution:** Marketplace for community schemas + custom schema builder

**Features:**
- ✅ Schema marketplace (upvote, share, monetize)
- ✅ Visual schema builder (drag-and-drop)
- ✅ API mocks (not just databases - Stripe, Twilio, GitHub APIs)
- ✅ Kubernetes operator (deploy in clusters)
- ✅ Terraform provider (IaC support)
- ✅ Enterprise features (SSO, air-gapped, SLA)

**Success metrics:**
- $100K MRR
- 10,000 paying users
- 100+ community schemas

---

### Phase 4: Developer Platform (Year 3+) - $500K+ MRR

**Vision:** The de-facto standard for realistic demo data

**Features:**
- ✅ Multi-cloud support (AWS, GCP, Azure)
- ✅ GraphQL API (not just REST)
- ✅ Real-time data streaming (Kafka, Kinesis)
- ✅ AI-powered schema generation (describe in English → generates schema)
- ✅ White-label (enterprises can rebrand)
- ✅ Partnerships (Snowflake, Databricks, Airbyte integrations)

**Success metrics:**
- $500K+ MRR
- 100K+ free users
- 20K+ paying users
- Profitable or path to profitability

---

## 🚀 Go-To-Market Strategy

### Channel 1: Open Source (Days 1-90) - Foundation

**Strategy:** Launch on GitHub, build community

**Tactics:**
1. **GitHub:** Polished README, examples, documentation
2. **Hacker News:** Launch post "Show HN: Realistic demo data in 30 seconds"
3. **Reddit:** Post to r/programming, r/datascience, r/devops
4. **ProductHunt:** Launch day 30 (after refining based on feedback)
5. **Twitter/X:** Tweet thread, engage with dev community

**Goal:** 1,000 GitHub stars by Month 3

---

### Channel 2: Developer Content (Months 3-12) - Awareness

**Strategy:** SEO + educational content

**Content:**
1. **Blog posts:**
   - "The Problem with Faker (and how to fix it)"
   - "How to Create Realistic Test Data in 30 Seconds"
   - "Why Your Demo Data Sucks (And How to Fix It)"

2. **Video tutorials:**
   - "Spin up a fintech database in 30 seconds"
   - "Build a demo app with realistic data"
   - "CI/CD with SourceBox"

3. **Documentation:**
   - Schema library
   - API reference
   - Integration guides

**Goal:** 10K monthly website visitors by Month 12

---

### Channel 3: Product-Led Growth (Months 6-18) - Conversion

**Strategy:** Free → paid conversion

**Tactics:**
1. **In-CLI prompts:** "Need cloud hosting? Try `sourcebox cloud` (14-day trial)"
2. **Usage-based prompts:** "You've used 10 schemas this month. Upgrade for unlimited?"
3. **Team features:** "Invite teammates to share schemas (Team plan)"
4. **Email onboarding:** Drip campaign for new users

**Goal:** 5% free → paid conversion by Month 12

---

### Channel 4: Partnerships (Year 2+) - Scale

**Strategy:** Partner with data companies

**Partners:**
1. **Airbyte:** "Official demo data provider"
2. **Snowflake:** Integration, marketplace listing
3. **Databricks:** Integration, co-marketing
4. **Docker:** Featured on Docker Hub
5. **GitHub:** GitHub Actions marketplace

**Goal:** 40%+ signups from partnerships by Year 2

---

## 🏁 Success Metrics

### North Star Metric

**Weekly active developers using SourceBox** (target: 10K+ by Month 12)

### Key Metrics (by Phase)

**Phase 1 (MVP - Months 1-6):**
- ✅ 1,000 GitHub stars
- ✅ 10,000 Docker Hub pulls
- ✅ 100 daily CLI installs
- ✅ NPS 50+ (product/market fit)

**Phase 2 (Cloud Launch - Months 7-12):**
- ✅ $10K MRR
- ✅ 1,000 paying users
- ✅ 5% free → paid conversion
- ✅ <5% monthly churn

**Phase 3 (Scale - Year 2):**
- ✅ $100K MRR
- ✅ 10,000 paying users
- ✅ 50+ community schemas
- ✅ Profitable or path to profitability

**Phase 4 (Platform - Year 3+):**
- ✅ $500K+ MRR
- ✅ 100K+ free users
- ✅ 20K+ paying users
- ✅ 130%+ net revenue retention

---

## 🛠️ Technical Strategy

### Core Architecture

**CLI Tool:**
- **Language:** Rust (performance) or Go (ecosystem + familiarity)
- **Distribution:** npm, Homebrew, cargo
- **Commands:** `seed`, `generate`, `reset`, `snapshot`, `cloud`
- **Config:** YAML files (`.sourcebox.yml`)

**Docker Images:**
- **Base:** Official MySQL, Postgres, MongoDB images
- **Layer:** Pre-seeded data + init scripts
- **Registry:** Docker Hub (free tier)
- **Naming:** `sourcebox/mysql-fintech`, `sourcebox/postgres-healthcare`

**Data Generation:**
- **Faker:** Names, addresses, emails (generic)
- **Custom logic:** Industry-specific patterns (loan amounts, patient visits)
- **Templates:** JSON/YAML schema definitions
- **Distributions:** Realistic (e.g., credit scores follow normal curve)

**Cloud Platform (Phase 2):**
- **Hosting:** AWS Fargate or Fly.io (containers)
- **Database:** Supabase (user management, API)
- **Storage:** S3 (seed data, snapshots)
- **API:** REST + GraphQL
- **Frontend:** Next.js (dashboard)

### Tech Stack

| Layer | Technology | Why |
|-------|------------|-----|
| **CLI** | Rust or Go | Fast, single binary, cross-platform |
| **Docker** | Docker Hub | Standard for developers |
| **Cloud** | AWS Fargate / Fly.io | Serverless containers, low cost |
| **Database** | Supabase | Postgres + auth + API out of box |
| **Frontend** | Next.js + Tailwind | Fast, SEO-friendly |
| **Docs** | Docusaurus | Standard for dev tools |
| **CI/CD** | GitHub Actions | Free for open source |

### MVP Scope (12 Weeks)

**Week 1-2: Validate**
- Interview 20 developers (backend, data engineers)
- Validate: "Would you use a CLI tool for demo data?"
- Validate: "Would you pay $20/mo for cloud version?"
- Goal: 15/20 say "yes" to CLI, 10/20 say "yes" to paid

**Week 3-8: Build Core**
- CLI tool (Rust/Go)
  - `sourcebox seed mysql --schema=fintech-loans`
  - `sourcebox seed postgres --schema=healthcare-patients`
- Docker images
  - MySQL + Postgres with 3 schemas each
- Data generators
  - Fintech, healthcare, retail schemas
  - 1,000 records each
- Documentation site

**Week 9-10: Polish**
- GitHub README + examples
- Docker Compose templates
- Video tutorial (3 min)
- Website landing page

**Week 11-12: Launch**
- Hacker News post
- Reddit posts (r/programming, r/devops)
- ProductHunt launch
- Twitter thread

**Success criteria:**
- ✅ 100 GitHub stars (Week 12)
- ✅ 1,000 Docker Hub pulls (Week 12)
- ✅ 20+ community feedback threads
- ✅ NPS 40+ (from early users)

---

## 💰 Funding Strategy

### Option A: Bootstrap (Recommended)

**Rationale:** Can build MVP with $0 (nights/weekends)

**Timeline:**
- **Months 1-3:** Build MVP, launch open source
- **Months 4-6:** Grow GitHub community (1K stars)
- **Months 7-9:** Launch cloud version ($20/mo freemium)
- **Months 10-12:** Hit $10K MRR

**Outcome:** Profitable by Month 12, no dilution

---

### Option B: Pre-Seed ($500K)

**Use of funds:**
- $200K: Founder salary (12 months)
- $150K: 1 engineer (12 months)
- $100K: Infrastructure + marketing
- $50K: Legal, accounting, misc

**Milestones:**
- **Month 3:** MVP launched, 1K GitHub stars
- **Month 6:** Cloud version launched, $5K MRR
- **Month 12:** $50K MRR
- **Month 18:** Raise seed round ($2-3M)

**Investors to target:**
- Y Combinator (dev tools focus)
- Sequoia Arc (early-stage dev tools)
- Hustle Fund (PLG SaaS)
- Individual angels (ex-founders of Retool, Supabase, Vercel)

---

## ⚠️ Key Risks & Mitigations

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| **Adoption risk** (devs don't use CLI) | High | Medium | Launch open source first, validate demand before building cloud |
| **Monetization risk** (free users don't convert) | High | Medium | Start with generous free tier, charge for cloud/teams (clear value prop) |
| **Competitive risk** (Faker adds verticalization) | Medium | Low | Network effects (50+ schemas), move fast (12-month lead) |
| **Technical risk** (data generation is hard) | Medium | Low | Start simple (Faker + templates), iterate based on feedback |
| **Legal risk** (employment IP claims) | High | Low | Develop on personal computer, consult lawyer, use MIT license |

---

## 🎯 The Ask (If Raising)

**Seeking:** $500K pre-seed at $3-5M post-money valuation

**Why now?**
1. **Market timing:** Developers frustrated with generic Faker
2. **Infrastructure ready:** Docker, serverless make this feasible at low cost
3. **TAM expansion:** $290M+ opportunity (vs niche tools)
4. **Founder/market fit:** Data engineer building for data engineers

**Milestones (18 months):**
- Month 3: 1K GitHub stars, MVP launched
- Month 6: $5K MRR (cloud version)
- Month 12: $50K MRR, profitable unit economics
- Month 18: $150K MRR, raise seed round

---

## 📅 Timeline

**Now - Month 3:** Validate & Build MVP (open source)
**Month 3-6:** Launch open source, grow community
**Month 6-9:** Launch cloud version (freemium)
**Month 9-12:** Scale to $10K MRR
**Month 12-18:** Scale to $50K MRR
**Month 18+:** Raise seed, hire team, scale to $500K+ MRR

---

## 🏆 Why Now?

1. **Developer productivity focus:** Companies invest in dev tools post-COVID
2. **Cloud infrastructure commoditized:** Docker, serverless make this cheap to build
3. **PLG proven:** Postman, Docker, Vercel show devs will pay for great tools
4. **Data literacy rising:** More devs work with data → need better demo data
5. **AI boom:** More AI demos → need realistic training data

---

## 👤 Founder Profile

**Ideal Founder:** Data engineer / backend developer with 3-5 years experience

**Unfair Advantages:**
- ✅ Lives the problem daily (uses Faker, frustrated with it)
- ✅ Knows the ecosystem (Docker, databases, dev tools)
- ✅ Understands verticalization (fintech, healthcare data patterns)
- ✅ Technical chops (can build MVP alone)
- ✅ Community credibility (dev building for devs)

**Skills Needed:**
- Technical: Backend development (Rust/Go), databases, Docker
- Product: What to build, what to skip (MVP discipline)
- Marketing: Open source community building, content creation
- Grit: Nights/weekends for 6-9 months

---

## 📚 References & Research

**Primary Research:**
- 20+ developer interviews (backend devs, data engineers)
- r/programming, r/datascience pain point analysis
- HN threads on demo data problems

**Market Research:**
- Stack Overflow Survey 2024 (dev tool usage)
- State of Developer Tools 2024 (PLG trends)
- Docker Hub stats (adoption patterns)

**Competitive Analysis:**
- Faker.js (generic, not verticalized)
- Mockaroo (web-based, not CLI)
- Manual SQL scripts (tedious, doesn't scale)
- Production dumps (security risk, compliance issues)

---

## 🚢 Ship Fast Philosophy

**Core Principles:**
1. **Ship MVP in 12 weeks** (not 12 months)
2. **Open source first** (validate demand before monetizing)
3. **Start with 3 schemas** (fintech, healthcare, retail - prove it works)
4. **CLI-first** (developers love terminals, not web UIs)
5. **Boring tech** (Rust/Go, Docker, Postgres - proven > shiny)

**What to skip in MVP:**
- ❌ Cloud version (Phase 2)
- ❌ Custom schema builder (Phase 3)
- ❌ API mocks (Phase 3)
- ❌ Marketplace (Phase 3)
- ❌ Perfect documentation (iterate based on feedback)

**What's critical in MVP:**
- ✅ CLI works flawlessly (`sourcebox seed` just works)
- ✅ Docker images pre-built (30-second spin-up)
- ✅ Data looks realistic (passes "sniff test")
- ✅ 3 industries (fintech, healthcare, retail)
- ✅ 2 databases (MySQL, Postgres)

---

**Last Updated:** 2025-01-14
**Version:** 2.0 (Developer Platform Vision)
**Status:** Pre-Launch / Greenfield

**Legal Note:** This project is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information.
