# SourceBox: Year 2 Vision (Months 13-24)

> **Vision**: From CLI Tool → Platform + Marketplace
> **Goal**: $10K MRR → $100K MRR, 10,000 paying users, community-driven schema marketplace
> **Status**: Year 2 Planning
> **Last Updated**: 2025-01-14

---

## 📋 Executive Summary

**Year 2 Objective**: Transform SourceBox from a CLI tool with cloud hosting into a platform with a thriving schema marketplace, API ecosystem, and enterprise features.

**Strategic Shift**: Year 1 = "Build → Validate → Monetize" (single product, freemium SaaS)
**Year 2 Strategy**: "Platform → Marketplace → Enterprise" (ecosystem, community, B2B sales)

**Phases**:
- **Q1 (Months 13-15)**: Schema Marketplace Launch (buy/sell schemas, revenue share)
- **Q2 (Months 16-18)**: API Platform + Integrations (API access, CI/CD, webhooks)
- **Q3 (Months 19-21)**: Enterprise Features (SSO, custom schemas, white-label)
- **Q4 (Months 22-24)**: Scale to $100K MRR (10,000 paying users, 5-person team)

**Success Metrics by Month 24**:
- ✅ **MRR**: $100K (10x from Year 1)
- ✅ **Paying users**: 10,000 (100x from Year 1)
- ✅ **Marketplace schemas**: 100+ (community-contributed)
- ✅ **Marketplace revenue**: $50K/year (20% commission)
- ✅ **Enterprise customers**: 10 ($20K/year each)
- ✅ **Team size**: 5 (founder + 4 FTEs)

---

## 🚀 Q1 (Months 13-15): Schema Marketplace Launch

**Goal**: Launch community-driven schema marketplace, enable creators to monetize schemas.

**Vision**: "GitHub Marketplace for demo data" - anyone can publish schemas, users pay per download or subscription.

---

### Month 13: Marketplace Foundation

**Week 49-50: Schema Publishing System**
- [ ] Design marketplace schema format: JSON + metadata (name, description, price, author, screenshots)
- [ ] Implement schema upload: Web UI (drag-and-drop JSON), validation, preview
- [ ] Implement schema moderation: Review queue (founder approves before publishing)
- [ ] Create schema detail page: Name, description, price, author, screenshots, reviews, download button
- [ ] Test internally: Upload 5 test schemas (SaaS metrics, logistics, education, insurance, real estate)

**Deliverable**: Schema publishing system working, 5 test schemas ready

**Week 51-52: Payment Processing**
- [ ] Implement Stripe Connect: Onboard creators (bank account, tax info)
- [ ] Implement revenue share: 80/20 split (creator gets 80%, SourceBox gets 20%)
- [ ] Implement payouts: Monthly payouts (automatic, $50 minimum)
- [ ] Implement purchase flow: Buy schema → Stripe checkout → download link
- [ ] Test end-to-end: Creator uploads schema, user buys, creator receives payout

**Deliverable**: Payment processing working, revenue share implemented

**Success Metrics**:
- Marketplace: Live at marketplace.sourcebox.dev
- Schemas published: 5 (test schemas)
- Payment flow: 100% success rate

---

### Month 14: Creator Recruitment

**Week 53-54: Onboard First 10 Creators**
- [ ] Identify 10 potential creators: Active GitHub contributors, Discord community members, data engineers
- [ ] Pitch marketplace: "Monetize your schemas, earn 80% revenue share"
- [ ] Offer early adopter incentives: Free Pro subscription (3 months), featured on homepage
- [ ] Help create schemas: 1-on-1 calls, schema template, examples, feedback
- [ ] Launch creators: Publish first 10 schemas (SaaS, logistics, education, insurance, real estate, manufacturing, energy, telecommunications, agriculture, hospitality)

**Deliverable**: 10 creators onboarded, 10 schemas published (total 15 including test schemas)

**Week 55-56: Marketing Push**
- [ ] Announce marketplace launch: Blog post, Hacker News (Show HN), Reddit r/SaaS, Product Hunt
- [ ] Highlight top creators: "Meet the creators building the SourceBox marketplace"
- [ ] Offer launch discount: 20% off all schemas for first month
- [ ] Track metrics: Schemas downloaded, revenue generated, creator payouts

**Deliverable**: Marketplace announced, 100+ downloads, $500 marketplace revenue

**Success Metrics**:
- Creators: 10 (active, published schemas)
- Schemas: 15 (5 test + 10 creator schemas)
- Downloads: 100+ (first month)
- Marketplace revenue: $500 (20% commission = $100 to SourceBox)

---

### Month 15: Marketplace Growth

**Week 57-58: Schema Discovery**
- [ ] Add search: Filter by industry, database, price, rating, downloads
- [ ] Add categories: Fintech, healthcare, retail, SaaS, e-commerce, logistics, etc.
- [ ] Add sorting: Most popular, highest rated, newest, price (low to high)
- [ ] Add reviews: 5-star rating, written reviews (moderated)
- [ ] Add featured schemas: Hand-picked by SourceBox team

**Deliverable**: Schema discovery improved, search + categories + sorting working

**Week 59-60: Creator Tools**
- [ ] Add creator dashboard: Revenue, downloads, reviews, analytics
- [ ] Add schema analytics: Downloads per day, revenue per schema, top users
- [ ] Add schema updates: Push new version, notify users (optional auto-update)
- [ ] Add affiliate program: Creators earn 10% for referring other creators
- [ ] Celebrate first $1K payout: Highlight top creator on blog, Twitter, Discord

**Deliverable**: Creator dashboard live, analytics working, first $1K payout sent

**Success Metrics**:
- Creators: 25 (10 from Month 14 + 15 new)
- Schemas: 50 (community-contributed)
- Downloads: 500+ (cumulative)
- Marketplace revenue: $5K (20% commission = $1K to SourceBox)
- Top creator payout: $1K (first month)

---

## ☁️ Q2 (Months 16-18): API Platform + Integrations

**Goal**: Build API platform for CI/CD, enable integrations with popular dev tools.

**Vision**: "SourceBox API as a service" - developers can automate demo data seeding in CI/CD pipelines.

---

### Month 16: API Platform

**Week 61-62: REST API**
- [ ] Design API: POST /api/v1/seed, POST /api/v1/environments, GET /api/v1/schemas
- [ ] Implement authentication: API keys (generate, rotate, revoke)
- [ ] Implement rate limiting: 100 requests/hour (free), 1,000 requests/hour (Pro), unlimited (Team/Enterprise)
- [ ] Implement webhooks: POST to customer URL (environment ready, environment expired, seed completed)
- [ ] Write API documentation: OpenAPI spec, interactive docs (Swagger UI)

**Deliverable**: REST API live, API keys working, documentation published

**Week 63-64: CLI Integration**
- [ ] Add API mode to CLI: `sourcebox seed --api-key=xxx --cloud` (seeds cloud DB, not local)
- [ ] Add CI/CD examples: GitHub Actions, GitLab CI, CircleCI, Travis CI
- [ ] Create quickstart guides: "How to seed demo data in your CI/CD pipeline"
- [ ] Test end-to-end: Run CLI in GitHub Actions, verify cloud DB seeded

**Deliverable**: CLI API mode working, CI/CD examples published

**Success Metrics**:
- API endpoints: Live at api.sourcebox.dev
- API documentation: Published at docs.sourcebox.dev/api
- API usage: 1,000+ requests/day
- CI/CD integrations: 50+ (GitHub Actions, GitLab CI, etc.)

---

### Month 17: Integrations

**Week 65-66: Developer Tool Integrations**
- [ ] Stripe integration: Mock Stripe API (payments, subscriptions, invoices)
- [ ] Twilio integration: Mock Twilio API (SMS, voice, webhooks)
- [ ] SendGrid integration: Mock SendGrid API (email sending, templates)
- [ ] Auth0 integration: Mock Auth0 API (user auth, SSO)
- [ ] Publish on marketplaces: Stripe App Marketplace, Twilio Marketplace

**Deliverable**: 4 integrations live, published on 2 marketplaces

**Week 67-68: Data Pipeline Integrations**
- [ ] Airbyte integration: SourceBox as Airbyte source (replicate demo data)
- [ ] dbt integration: SourceBox seed data for dbt testing
- [ ] Metabase integration: Pre-built dashboards with SourceBox data
- [ ] Fivetran integration: SourceBox as Fivetran source (demo pipelines)
  - **Legal Note**: Only integrate with publicly available APIs, no proprietary info
- [ ] Announce integrations: Blog post, partner announcements

**Deliverable**: 4 data pipeline integrations live, partner announcements published

**Success Metrics**:
- Integrations: 8 (Stripe, Twilio, SendGrid, Auth0, Airbyte, dbt, Metabase, Fivetran)
- Integration usage: 500+ (combined)
- Partner referrals: 100+ (from Stripe/Twilio marketplaces)

---

### Month 18: Developer Experience

**Week 69-70: SDKs & Libraries**
- [ ] Node.js SDK: `npm install @sourcebox/sdk` (seed data from Node.js apps)
- [ ] Python SDK: `pip install sourcebox` (seed data from Python apps)
- [ ] Go SDK: `go get github.com/sourcebox/sdk` (seed data from Go apps)
- [ ] Examples: Next.js, Django, Rails, Express, Flask
- [ ] Publish SDKs: npm, PyPI, Go modules

**Deliverable**: 3 SDKs published, examples live

**Week 71-72: Developer Portal**
- [ ] Create developer portal: portal.sourcebox.dev
- [ ] Add interactive docs: Try API in browser (Swagger UI, Postman-like)
- [ ] Add code examples: Copy-paste snippets (Node.js, Python, Go, curl)
- [ ] Add API playground: Test API without writing code
- [ ] Add API changelog: Track breaking changes, deprecations, new features

**Deliverable**: Developer portal live, interactive docs working

**Success Metrics**:
- SDKs: 3 (Node.js, Python, Go)
- SDK installs: 1,000+ (combined, first month)
- Developer portal: Live at portal.sourcebox.dev
- API playground: 500+ test requests/day

---

## 🏢 Q3 (Months 19-21): Enterprise Features

**Goal**: Add enterprise features (SSO, custom schemas, white-label), close first 10 enterprise deals ($20K/year each).

**Vision**: "SourceBox for teams" - large companies use SourceBox across engineering, sales, QA, DevRel.

---

### Month 19: Enterprise Infrastructure

**Week 73-74: SSO & Security**
- [ ] Implement SAML SSO: Okta, Azure AD, Google Workspace
- [ ] Implement SCIM provisioning: Auto-provision users, auto-deprovision on offboarding
- [ ] Implement audit logs: Track all user actions (who, what, when, where)
- [ ] Implement data isolation: Dedicated VPC per enterprise (optional)
- [ ] Implement compliance: SOC 2 Type I prep (policies, training, vendor assessment)

**Deliverable**: SAML SSO working, audit logs live, SOC 2 Type I in progress

**Week 75-76: Custom Schemas**
- [ ] Offer professional services: "We'll build custom schemas for your industry ($5K-$20K)"
- [ ] Hire 1 contract data engineer: Build custom schemas (20-40 hours/month)
- [ ] Create 3 enterprise schemas: Financial services, healthcare (HIPAA-compliant), government
- [ ] Deliver to customers: Custom schema + documentation + support
- [ ] Upsell to existing customers: "We noticed you're in fintech. Want a custom schema?"

**Deliverable**: Professional services live, 3 custom schemas delivered

**Success Metrics**:
- SAML SSO: Working with Okta, Azure AD, Google
- Audit logs: 100% of actions logged
- Custom schemas: 3 delivered ($15K revenue)
- SOC 2 Type I: In progress (6-month timeline)

---

### Month 20: Enterprise Sales

**Week 77-78: Sales Outreach**
- [ ] Identify 50 enterprise targets: Series B-C SaaS, data integration, workflow automation
- [ ] Cold email outreach: "SourceBox can save your engineering team 5-10 hours/week"
- [ ] LinkedIn outreach: Connect with VPs of Engineering, CTOs, Engineering Managers
- [ ] Offer 1-month free trial: "Try SourceBox Team free for 1 month, no credit card"
- [ ] Follow up: 3-day, 7-day, 14-day emails

**Deliverable**: 50 outreach emails sent, 10 trials started, 3 enterprise deals closed

**Week 79-80: Sales Collateral**
- [ ] Create enterprise deck: Problem, solution, ROI, case studies, pricing
- [ ] Create ROI calculator: "Your team spends X hours/week on demo data. SourceBox saves Y hours, worth $Z."
- [ ] Create case studies: 3 customer case studies (with permission)
- [ ] Create comparison table: SourceBox vs Faker.js vs Mockaroo vs manual SQL
- [ ] Publish on website: /enterprise page

**Deliverable**: Enterprise deck created, case studies published

**Success Metrics**:
- Enterprise trials: 10 (started)
- Enterprise deals: 5 (closed, $20K/year each = $100K ARR)
- Sales pipeline: $500K (25 opportunities × $20K average deal size)
- Sales cycle: 60 days (from first touch to close)

---

### Month 21: Enterprise Growth

**Week 81-82: White-Label**
- [ ] Implement white-label: Rebrand as "YourCo Demo Platform powered by SourceBox"
- [ ] Custom domain: demo.yourco.com (CNAME to SourceBox)
- [ ] Custom branding: Logo, colors, fonts, favicon
- [ ] Custom email: no-reply@yourco.com (instead of @sourcebox.dev)
- [ ] Offer to enterprise customers: "Rebrand SourceBox as your internal tool"

**Deliverable**: White-label working, 2 enterprise customers using

**Week 83-84: Enterprise Success**
- [ ] Hire 1 customer success manager: Onboard enterprise customers, training, support
- [ ] Create onboarding program: 1-hour kickoff, 1-week check-in, 1-month review
- [ ] Create training materials: Video tutorials, documentation, best practices
- [ ] Implement NPS surveys: Quarterly NPS surveys (enterprise customers)
- [ ] Celebrate 10 enterprise deals: Blog post, thank customers, share milestones

**Deliverable**: Customer success manager hired, onboarding program live, 10 enterprise deals closed

**Success Metrics**:
- Enterprise customers: 10 (total, $20K/year each = $200K ARR)
- White-label: 2 customers using
- Enterprise NPS: 70+ (strong product-market fit)
- Enterprise churn: < 5% (high retention)

---

## 📈 Q4 (Months 22-24): Scale to $100K MRR

**Goal**: Scale from $10K → $100K MRR, grow team to 5 people, prepare for Series A.

**Vision**: "SourceBox as a category leader" - the standard tool for demo data, used by 10,000+ developers.

---

### Month 22: Team Building

**Week 85-86: Hire Full-Time Team**
- [ ] Hire 1 full-stack engineer: Go + React, build features, fix bugs ($120K/year)
- [ ] Hire 1 DevOps engineer: AWS, Kubernetes, monitoring, scaling ($140K/year)
- [ ] Hire 1 product manager: Roadmap, prioritization, user research ($130K/year)
- [ ] Hire 1 customer success manager: Onboarding, support, retention ($100K/year)
- [ ] Total team: 5 (founder + 4 FTEs, $490K/year burn)

**Deliverable**: 4 FTEs hired, onboarding complete

**Week 87-88: Team Processes**
- [ ] Set up project management: Linear (tasks, sprints, roadmap)
- [ ] Set up communication: Slack (team channels), Zoom (daily standups)
- [ ] Set up documentation: Notion (internal wiki, onboarding docs)
- [ ] Set up deployment: GitHub Actions (CI/CD, auto-deploy to staging/prod)
- [ ] Set up monitoring: Datadog (APM, logs, metrics, alerts)

**Deliverable**: Team processes in place, tools set up

**Success Metrics**:
- Team size: 5 (founder + 4 FTEs)
- Burn rate: $50K/month ($490K/year salaries + $110K/year infra/tools)
- Runway: 20+ months (at $100K MRR, $50K profit/month after burn)

---

### Month 23: Growth Acceleration

**Week 89-90: Marketing Campaigns**
- [ ] Launch referral program: "Refer 5 friends, get 1 year free Pro"
- [ ] Launch affiliate program: "Earn 20% recurring commission for each referral"
- [ ] Launch podcast sponsorships: Sponsor 5 developer podcasts ($2K-$5K/episode)
- [ ] Launch conference sponsorships: Sponsor 3 developer conferences ($10K-$20K/conference)
- [ ] Track ROI: Cost per acquisition (CPA), lifetime value (LTV), LTV/CAC ratio

**Deliverable**: Marketing campaigns launched, CPA < $100, LTV/CAC > 3

**Week 91-92: Product Expansion**
- [ ] Add MongoDB support: CLI + Docker + cloud hosting
- [ ] Add GraphQL API: Complement REST API (for complex queries)
- [ ] Add real-time collaboration: Share environments with teammates, live updates
- [ ] Add advanced analytics: Track which features used, win rates, ROI
- [ ] Ship 10 new schemas: Community-contributed + built-in

**Deliverable**: MongoDB support live, GraphQL API live, 10 new schemas shipped

**Success Metrics**:
- MRR: $50K (5x from $10K in Month 12)
- Paying users: 5,000 (50x from 100 in Month 12)
- Marketplace schemas: 100+ (community-contributed)
- Enterprise customers: 15 (5 new in this month)

---

### Month 24: Scale to $100K MRR

**Week 93-94: Final Push**
- [ ] Optimize conversion: A/B test pricing, onboarding, upgrade prompts
- [ ] Optimize retention: Reduce churn (5% → 3%), improve NPS (60 → 70)
- [ ] Optimize infrastructure: Reduce AWS costs (Spot instances, reserved instances)
- [ ] Close 10 more enterprise deals: $200K ARR (10 × $20K/year)
- [ ] Celebrate $100K MRR: Blog post, press release, team celebration

**Deliverable**: $100K MRR reached, team celebrates

**Week 95-96: Series A Prep**
- [ ] Create pitch deck: Vision, traction, team, financials, roadmap
- [ ] Create financial model: Revenue, burn, runway, unit economics
- [ ] Create investor list: 20 Seed/Series A VCs (focus on B2B SaaS, dev tools)
- [ ] Schedule meetings: 10 intro calls (via warm intros)
- [ ] Decide: Raise Series A ($2M-$5M) or bootstrap to profitability?

**Deliverable**: Series A deck created, investor meetings scheduled

**Success Metrics**:
- MRR: $100K (from $10K in 12 months = 10x growth)
- ARR: $1.2M (annual recurring revenue)
- Paying users: 10,000 (from 100 in 12 months = 100x growth)
- Enterprise customers: 25 ($500K ARR)
- Marketplace revenue: $50K/year (20% commission)
- Team size: 5 (founder + 4 FTEs)
- Gross margin: 85% ($100K revenue - $15K infra/tools)
- Net margin: 35% ($100K revenue - $50K burn - $15K infra)

---

## 📊 Revenue Model Breakdown (Month 24)

### Revenue Streams

| Stream | Users | Price | MRR | % of Total |
|--------|-------|-------|-----|------------|
| **Free** | 50,000 | $0 | $0 | 0% |
| **Pro** | 8,000 | $20/mo | $160K | 64% |
| **Team** | 1,500 | $200/mo | $300K | 30% |
| **Enterprise** | 25 | $20K/year | $42K/mo | 17% |
| **Marketplace** | - | 20% commission | $4K/mo | 1% |
| **Total** | 59,525 | - | **$506K/mo** | **100%** |

**Note**: MRR from monthly plans ($460K) + amortized annual plans ($42K) + marketplace ($4K) = $506K/mo.

### Unit Economics (Month 24)

| Metric | Value | Target |
|--------|-------|--------|
| Average revenue per user (ARPU) | $8.50/mo | $10/mo |
| Gross margin | 85% ($100K - $15K infra) | 80%+ |
| Net margin | 35% ($100K - $50K burn - $15K) | 30%+ |
| Customer acquisition cost (CAC) | $150 | < $200 |
| Lifetime value (LTV) | $500 (blended) | > $450 |
| LTV/CAC ratio | 3.3x | > 3x |
| Monthly churn | 3% | < 5% |
| Net revenue retention | 120% | > 110% |

### Cost Structure (Month 24)

| Category | Cost/Month | % of Revenue |
|----------|------------|--------------|
| Salaries (4 FTEs) | $40K | 40% |
| Infrastructure (AWS, Vercel, Supabase) | $10K | 10% |
| Marketing (ads, content, events) | $5K | 5% |
| Tools (Datadog, Stripe, Linear, etc.) | $2K | 2% |
| Miscellaneous (legal, accounting) | $3K | 3% |
| **Total** | **$60K** | **60%** |

**Profit**: $100K revenue - $60K cost = **$40K/month** ($480K/year)

---

## 🎯 Key Assumptions & Risks

### Assumptions

1. **Marketplace takes off**: 100+ community schemas by Month 24
2. **Enterprise deals close**: 25 enterprises at $20K/year each
3. **Team execution**: 4 FTEs deliver features on time
4. **Churn stays low**: < 5% monthly churn (85%+ retention)
5. **Free → Paid conversion**: 10% (50,000 free → 5,000 paid)

### Risks & Mitigations

**Risk 1: Marketplace fails (< 20 schemas)**
- **Impact**: High (no community growth, no marketplace revenue)
- **Probability**: Medium
- **Mitigation**: Pay creators upfront ($500-$1K per schema), highlight top creators, improve discoverability

**Risk 2: Enterprise sales too slow (< 10 deals)**
- **Impact**: High (can't hit $100K MRR)
- **Probability**: Medium
- **Mitigation**: Hire sales rep (Month 18), offer free trials, create case studies, lower price to $10K/year

**Risk 3: Team can't execute (features delayed)**
- **Impact**: High (lose competitive advantage)
- **Probability**: Low
- **Mitigation**: Hire experienced team, use Linear for prioritization, ship MVP versions first

**Risk 4: Burn too high (> $60K/month)**
- **Impact**: Medium (runway shortens)
- **Probability**: Low
- **Mitigation**: Optimize infra costs (Spot instances), delay hiring, cut marketing spend

**Risk 5: Competitor launches (Faker Pro, Mockaroo 2.0)**
- **Impact**: Medium (lose differentiation)
- **Probability**: Medium
- **Mitigation**: Ship fast (marketplace first), build moat (community schemas), focus on API/integrations

---

## 🏆 Success Milestones

- **$20K MRR** (Month 13-15): 🎂 Marketplace launched
- **$50K MRR** (Month 19-21): 🎉 First 10 enterprise deals
- **$75K MRR** (Month 22-23): 🚀 Team of 5 hired
- **$100K MRR** (Month 24): 🏆 10x growth from Year 1, Series A decision

---

## 📚 Strategic Priorities

### Q1 (Months 13-15): Marketplace
**Why**: Build community, enable creators, expand schema library from 10 → 100+.
**Success**: 25 creators, 50 schemas, $5K marketplace revenue.

### Q2 (Months 16-18): API Platform
**Why**: Enable CI/CD, integrate with dev tools, expand TAM (developers + data engineers + SEs).
**Success**: 8 integrations, 1,000+ API requests/day, 50+ CI/CD users.

### Q3 (Months 19-21): Enterprise
**Why**: Higher ACV ($20K/year vs $240/year), lower churn, strategic logos.
**Success**: 10 enterprise deals, $200K ARR, SOC 2 Type I in progress.

### Q4 (Months 22-24): Scale
**Why**: Hire team, accelerate growth, prepare for Series A or profitability.
**Success**: $100K MRR, 10,000 paying users, 5-person team.

---

## 🚧 Year 2 vs Year 1 Comparison

| Metric | Year 1 (Month 12) | Year 2 (Month 24) | Growth |
|--------|-------------------|-------------------|--------|
| **MRR** | $10K | $100K | **10x** |
| **Paying users** | 100 | 10,000 | **100x** |
| **Schemas** | 6 (built-in) | 100+ (marketplace) | **16x** |
| **Integrations** | 0 | 8 (Stripe, Twilio, etc.) | **∞** |
| **Enterprise customers** | 0 | 25 ($500K ARR) | **∞** |
| **Team size** | 1 (founder) | 5 (founder + 4 FTEs) | **5x** |
| **Gross margin** | 92% | 85% | -7% |
| **Net margin** | 50% | 35% | -15% |

**Analysis**: Revenue 10x, team 5x, margins down (investing in growth). Net profit still 35% ($40K/month = $480K/year).

---

## 📝 Decision Framework for Year 2

When making decisions during Year 2, ask:

1. **Does this move us toward $100K MRR?** (Yes → prioritize, No → defer)
2. **Does this expand the platform/marketplace?** (Yes → prioritize, No → defer)
3. **Does this enable enterprise sales?** (Yes → prioritize, No → defer)
4. **Does this require > 2 engineers for > 1 month?** (Yes → break into phases, No → build)
5. **Does this violate legal constraints?** (Yes → reject, No → consider)

**Examples**:
- Launch marketplace: ✅ Yes (expands ecosystem, enables community)
- Add GraphQL API: ✅ Yes (expands platform, enables integrations)
- Build visual schema builder: ⚠️ Maybe (nice-to-have, defer to Q4)
- Add AI-generated schemas: ❌ No (not a priority, defer to Year 3)
- Integrate with Fivetran: ✅ Yes (but only with public APIs, no proprietary info)

---

## 🎉 Year 2 End Goal

**Vision Achieved**: SourceBox is the category leader for demo data.
- **Ecosystem**: 100+ community schemas, 25 creators earning $1K+/month
- **Platform**: 8 integrations, API used by 1,000+ developers
- **Enterprise**: 25 strategic customers ($500K ARR)
- **Team**: 5-person team shipping fast
- **Revenue**: $100K MRR ($1.2M ARR), 35% net margin ($480K profit/year)
- **Decision**: Raise Series A ($2M-$5M) or bootstrap to profitability?

---

## 📚 Resources

**Key Documents**:
- [BRIEF.md](./BRIEF.md) - Project vision, market opportunity, roadmap
- [PRD.md](./PRD.md) - Product requirements, features, specifications
- [MVP-ROADMAP.md](./MVP-ROADMAP.md) - 12-week implementation plan (Months 1-3)
- [YEAR-1-ROADMAP.md](./YEAR-1-ROADMAP.md) - Month-by-month plan (Months 1-12)
- [constitution-prompt.md](./constitution-prompt.md) - Project principles
- [docs/technical-summary.yaml](./docs/technical-summary.yaml) - Technical architecture

**Community**:
- GitHub: https://github.com/yourusername/sourcebox
- Discord: https://discord.gg/sourcebox
- Marketplace: https://marketplace.sourcebox.dev
- Developer Portal: https://portal.sourcebox.dev

---

## ⚖️ Legal Notice

**This project is developed independently** on personal equipment, outside of work hours, with no use of employer resources or proprietary information.

All references to companies or products in this documentation are for illustrative purposes only. SourceBox is a standalone, open source project with no affiliation to any third party.

---

**Last Updated**: 2025-01-14
**Version**: 1.0
**Status**: Year 2 Planning
**Next Review**: End of Year 1 (Month 12)
