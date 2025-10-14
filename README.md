# SourceBox

> **Realistic, verticalized demo data for developers** - spin up production-like databases in 30 seconds.

[![GitHub stars](https://img.shields.io/github/stars/yourusername/sourcebox?style=social)](https://github.com/yourusername/sourcebox)
[![npm version](https://img.shields.io/npm/v/sourcebox)](https://www.npmjs.com/package/sourcebox)
[![Docker Pulls](https://img.shields.io/docker/pulls/sourcebox/mysql-fintech)](https://hub.docker.com/u/sourcebox)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Status:** Pre-Launch / MVP Development
**Version:** 0.1.0 (Alpha)
**Created:** 2025-01-14

---

## 🎯 What is SourceBox?

**The problem:** Developers waste 5-10 hours per week creating realistic demo data. Faker is too generic ("John Doe", "Acme Corp"), manual SQL is tedious, and production dumps are a security nightmare.

**The solution:** SourceBox provides **verticalized demo data** - pre-built schemas for fintech, healthcare, retail, and more. Spin up a MySQL database with 1,000 realistic loan records in 30 seconds.

**Think of it as:** `docker run mysql` but with realistic data already seeded.

---

## ⚡ Quick Start

### Option 1: CLI

```bash
# Install
npm install -g sourcebox
# or
brew install sourcebox

# Seed a database
sourcebox seed mysql --schema=fintech-loans --records=1000

# Output:
✅ Seeded MySQL database 'demo' with 1,000 loan records
✅ Tables: loans, borrowers, payments, credit_scores
✅ Time: 15 seconds

# Connect and query
mysql -u root -p demo
mysql> SELECT * FROM loans LIMIT 5;
```

### Option 2: Docker

```bash
# Run pre-seeded MySQL with fintech data
docker run -p 3306:3306 sourcebox/mysql-fintech:latest

# Or use Docker Compose
version: '3.8'
services:
  db:
    image: sourcebox/mysql-fintech:latest
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: demo
```

### Option 3: Docker Compose (Multi-Database)

```yaml
# docker-compose.yml
version: '3.8'
services:
  mysql-fintech:
    image: sourcebox/mysql-fintech:latest
    ports:
      - "3306:3306"

  postgres-healthcare:
    image: sourcebox/postgres-healthcare:latest
    ports:
      - "5432:5432"

  mongo-ecommerce:
    image: sourcebox/mongo-ecommerce:latest
    ports:
      - "27017:27017"
```

```bash
docker-compose up -d
# Now you have 3 databases with realistic industry data
```

---

## 📚 Available Schemas

| Industry | Database | Tables | Records | Use Case |
|----------|----------|--------|---------|----------|
| **Fintech** | MySQL, Postgres | loans, borrowers, payments, credit_scores | 1,000 | Lending, banking, payments demos |
| **Healthcare** | Postgres, MySQL | patients, visits, prescriptions, claims | 500 | Healthcare apps, HIPAA demos |
| **Retail** | MySQL, Postgres, MongoDB | products, orders, customers, inventory | 2,000 | E-commerce, POS systems |
| **SaaS** | Postgres, MySQL | users, subscriptions, usage, billing | 1,000 | SaaS metrics, analytics |
| **E-commerce** | MongoDB, Postgres | orders, products, customers, reviews | 5,000 | Marketplaces, shopping carts |
| **Logistics** | MySQL, Postgres | shipments, routes, drivers, warehouses | 1,500 | Supply chain, delivery |
| **Education** | Postgres, MySQL | students, courses, grades, enrollment | 800 | EdTech, LMS systems |
| **Real Estate** | MySQL, Postgres | properties, agents, transactions, listings | 600 | PropTech, MLS systems |
| **Insurance** | Postgres, MySQL | policies, claims, premiums, customers | 1,000 | InsurTech, underwriting |
| **Manufacturing** | MySQL, Postgres | products, orders, inventory, suppliers | 1,200 | ERP, supply chain |

**More schemas coming soon.** [Request a schema →](https://github.com/yourusername/sourcebox/issues/new?template=schema-request.md)

---

## 🔥 Why SourceBox?

### The Problem with Current Options

| Solution | Problem |
|----------|---------|
| **Faker.js** | Generic data ("John Doe", "test@test.com") - not industry-specific |
| **Manual SQL** | Tedious to write 1,000+ INSERT statements |
| **Production dumps** | Security risk, GDPR/HIPAA compliance nightmare |
| **Copy/paste tutorials** | Outdated, incomplete, doesn't fit your use case |

### The SourceBox Advantage

✅ **Verticalized** - Fintech data looks like fintech (loan amounts, credit scores)
✅ **Realistic** - Proper distributions, relationships, edge cases
✅ **Fast** - 30 seconds from `docker run` to querying data
✅ **Multiple databases** - MySQL, Postgres, MongoDB, SQLite
✅ **Local-first** - Runs on your laptop, no cloud required
✅ **Open source** - MIT license, community-driven

---

## 🛠️ Use Cases

### For Backend Developers

```bash
# Spin up local dev database with realistic data
sourcebox seed postgres --schema=saas-metrics --records=1000

# Now build your analytics dashboard with real-looking data
```

### For Data Engineers

```bash
# Test your ETL pipeline with realistic source data
docker run sourcebox/mysql-fintech:latest

# Connect your data pipeline (Airbyte, dbt, etc.)
# Test transformations with realistic edge cases
```

### For Sales Engineers

```bash
# Spin up demo environment for customer
docker-compose up -d  # Fintech database ready in 30 seconds

# Show live queries with realistic healthcare patient data
# Customize data to match prospect's industry
```

### For QA Engineers

```bash
# Seed test database with edge cases
sourcebox seed mysql --schema=fintech-loans --edge-cases=true

# Test your app against realistic data scenarios
# Negative credit scores, large loan amounts, late payments, etc.
```

### For DevRel / Solutions Architects

```bash
# Create tutorial with realistic data
docker run sourcebox/postgres-healthcare:latest

# Build demo app that looks production-ready
# Blog post, video tutorial, conference talk
```

---

## 📖 Documentation

**Quick Links:**
- [Installation Guide](https://docs.sourcebox.dev/installation)
- [Schema Library](https://docs.sourcebox.dev/schemas)
- [CLI Reference](https://docs.sourcebox.dev/cli)
- [Docker Images](https://docs.sourcebox.dev/docker)
- [Custom Schemas](https://docs.sourcebox.dev/custom-schemas)
- [Examples](https://github.com/yourusername/sourcebox/tree/main/examples)

**Core Docs:**
1. **[BRIEF.md](./BRIEF.md)** - Project vision, market opportunity, roadmap
2. **[PRD.md](./PRD.md)** - Product requirements, features, specifications
3. **[MVP-ROADMAP.md](./MVP-ROADMAP.md)** - 12-week implementation plan
4. **[constitution-prompt.md](./constitution-prompt.md)** - Project principles (for `/speckit.constitution`)

---

## 🚀 Roadmap

### Phase 1: CLI + Docker MVP (Months 1-6)

**Now** (In Development):
- [x] Project vision & docs
- [ ] CLI tool (Rust/Go)
- [ ] 3 schemas (fintech, healthcare, retail)
- [ ] 2 databases (MySQL, Postgres)
- [ ] Docker images on Docker Hub
- [ ] Documentation site
- [ ] Open source launch

**Success Metrics:**
- 1,000 GitHub stars
- 10,000 Docker Hub pulls
- 100 daily CLI installs

### Phase 2: Cloud SaaS (Months 7-12)

- [ ] Cloud-hosted databases (on-demand)
- [ ] API access for CI/CD
- [ ] 50+ schemas (expand library)
- [ ] Custom schema builder
- [ ] Freemium pricing ($0 → $20/mo → $200/mo)

**Success Metrics:**
- $10K MRR
- 1,000 paying users
- 5% free → paid conversion

### Phase 3: Platform + Marketplace (Year 2)

- [ ] Community schema marketplace
- [ ] Visual schema builder (drag-and-drop)
- [ ] API mocks (Stripe, Twilio, GitHub)
- [ ] Kubernetes operator
- [ ] Terraform provider
- [ ] Enterprise features (SSO, air-gapped)

**Success Metrics:**
- $100K MRR
- 10,000 paying users
- 100+ community schemas

[See full roadmap →](./BRIEF.md#-product-vision)

---

## 🤝 Contributing

SourceBox is **open source** (MIT license) and **community-driven**.

**Ways to contribute:**
1. **Use it & give feedback** - [Open an issue](https://github.com/yourusername/sourcebox/issues)
2. **Request schemas** - [Schema request template](https://github.com/yourusername/sourcebox/issues/new?template=schema-request.md)
3. **Contribute schemas** - [Schema contribution guide](./CONTRIBUTING.md)
4. **Improve docs** - [Documentation repo](https://github.com/yourusername/sourcebox-docs)
5. **Report bugs** - [Bug report template](https://github.com/yourusername/sourcebox/issues/new?template=bug-report.md)

**Development:**
```bash
# Clone repo
git clone https://github.com/yourusername/sourcebox.git
cd sourcebox

# Install dependencies (TBD - depends on Rust vs Go)
cargo build  # if Rust
# or
go build  # if Go

# Run tests
cargo test
# or
go test ./...

# Build Docker images
cd docker/mysql-fintech
docker build -t sourcebox/mysql-fintech:latest .
```

See [CONTRIBUTING.md](./CONTRIBUTING.md) for full guide.

---

## 📊 Success Metrics

### North Star Metric

**Weekly active developers using SourceBox** (target: 10K+ by Month 12)

### Current Status (Pre-Launch)

| Metric | Target (Month 6) | Current | Status |
|--------|------------------|---------|--------|
| **GitHub stars** | 1,000 | 0 | 🔴 Pre-launch |
| **Docker Hub pulls** | 10,000 | 0 | 🔴 Pre-launch |
| **Daily CLI installs** | 100 | 0 | 🔴 Pre-launch |
| **NPS** | 50+ | - | 🔴 Not measured |

[See full metrics →](./BRIEF.md#-success-metrics)

---

## 🧰 Tech Stack

| Layer | Technology | Why |
|-------|------------|-----|
| **CLI** | Rust or Go | Fast, single binary, cross-platform |
| **Docker** | Docker Hub | Standard for developers |
| **Data generation** | Faker + custom logic | Realistic distributions |
| **Schemas** | JSON/YAML templates | Easy to extend |
| **Docs** | Docusaurus | Standard for dev tools |

**Cloud version** (Phase 2):
- AWS Fargate / Fly.io (serverless containers)
- Supabase (Postgres + auth + API)
- Next.js + Tailwind (dashboard)

[See full tech design →](./docs/technical-summary.yaml)

---

## 💡 Philosophy

### 1. Developer-First

Built by developers, for developers. CLI-first, Docker-native, no unnecessary abstractions.

### 2. Verticalized, Not Generic

Fintech data should look like fintech. Healthcare should look like healthcare. No more "Employee 123".

### 3. Open Source, Community-Driven

Free forever. MIT license. Community schemas. Transparent roadmap.

### 4. Local-First, Cloud Optional

Runs entirely on your laptop. No cloud required. Cloud version (Phase 2) is optional for teams.

### 5. Ship Fast, Iterate

12-week MVP. Launch with 3 schemas, not 30. Iterate based on feedback.

---

## ⚖️ License

MIT License - see [LICENSE](./LICENSE) for details.

**What this means:**
- ✅ Use for any purpose (commercial, personal, internal)
- ✅ Modify and distribute
- ✅ No attribution required (but appreciated!)
- ⚠️ No warranty (use at your own risk)

---

## 🙏 Acknowledgments

**Inspiration:**
- [Faker.js](https://fakerjs.dev/) - Great for generic data, inspired verticalization need
- [Mockaroo](https://www.mockaroo.com/) - Web-based, inspired CLI approach
- [Docker](https://www.docker.com/) - Made containerized data possible

**Built with:**
- Spec-Kit development framework ([github.com/github/spec-kit](https://github.com/github/spec-kit))
- Claude Code AI assistant (specification-driven development)

---

## 📞 Contact & Community

**Questions? Ideas? Feedback?**

- 💬 [GitHub Discussions](https://github.com/yourusername/sourcebox/discussions)
- 🐛 [Report a Bug](https://github.com/yourusername/sourcebox/issues/new?template=bug-report.md)
- 💡 [Request a Feature](https://github.com/yourusername/sourcebox/issues/new?template=feature-request.md)
- 📧 Email: hello@sourcebox.dev
- 🐦 Twitter: [@sourcebox_dev](https://twitter.com/sourcebox_dev)
- 💬 Discord: [Join community](https://discord.gg/sourcebox)

**Follow development:**
- ⭐ Star this repo for updates
- 👀 Watch releases
- 📣 Follow [@sourcebox_dev](https://twitter.com/sourcebox_dev) on Twitter

---

## 🚨 Important Notes

### Legal

This project is developed **independently** on personal equipment, outside of work hours, with no use of employer resources or proprietary information.

### Status

**Pre-Launch** - This is a greenfield project currently in specification phase. CLI tool and Docker images are not yet available.

**Want to help?** [Join early access](https://github.com/yourusername/sourcebox/discussions/1) or [contribute](./CONTRIBUTING.md).

---

**Last Updated:** 2025-01-14
**Version:** 0.1.0 (Alpha)
**Status:** Pre-Launch / MVP Development

**Star this repo** to follow development →
