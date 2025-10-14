# Feature Specification Prompt: F018 - Retail E-commerce Schema Definition (JSON)

## Feature Metadata
- **Feature ID**: F018
- **Name**: Retail E-commerce Schema Definition (JSON)
- **Category**: Data Generation
- **Phase**: Week 8
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F007 (Schema spec), F008 (Parser)

## Constitutional Alignment
- ✅ **Verticalized > Generic**: Retail data must look like retail
- ✅ **Realism**: Real products, realistic prices, seasonal patterns
- ✅ **Quality Bar**: 10x more realistic than Faker

## User Story
**US-MVP-012**: "As an e-commerce developer, I want realistic product, customer, order, and order item data with real product names and realistic purchase patterns."

## Solution Overview
Create `schemas/retail-ecommerce.json` with 4 tables: products (500), customers (1000), orders (2000), order_items (5000). Use real product names, realistic price distributions, seasonal order patterns.

## Acceptance Criteria
- Schema file: `schemas/retail-ecommerce.json`
- 4 tables with proper FK relationships
- Real products ('iPhone 15 Pro', 'Nike Air Max')
- Realistic prices (log-normal: $5-$2000)
- Order status: weighted (pending=10%, shipped=40%, delivered=40%, returned=10%)
- Seasonal patterns for order dates
- Total: 8,500 records

## Related Constitution Sections
- **Verticalized > Generic (Principle I)**
- **Realism**: Real products and pricing patterns
- **Quality Bar**: 10x more realistic than Faker
