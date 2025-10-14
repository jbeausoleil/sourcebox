# Feature Specification Prompt: F019 - Custom Retail Data Generators

## Feature Metadata
- **Feature ID**: F019
- **Name**: Custom Retail Data Generators
- **Category**: Data Generation
- **Phase**: Week 8
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (4 days)
- **Dependencies**: F011 (Base generators), F018 (Retail schema)

## Constitutional Alignment
- ✅ **Verticalized > Generic**: Retail data must use real products
- ✅ **Realism**: Realistic price distributions, categories
- ✅ **TDD Required**: Core generators need tests

## User Story
**US-MVP-013**: "As a data generation engine, I need generators for product names, prices, categories, and order statuses with realistic distributions."

## Solution Overview
Implement `pkg/generators/retail.go` with: ProductName() (real products), Price() (log-normal by category), OrderStatus() (weighted), Category() (realistic categories), seasonal order date patterns.

## Acceptance Criteria
- ProductName: real products ('iPhone 15 Pro', 'Nike Air Max')
- Price: log-normal ($5-$2000) varying by category
- OrderStatus: weighted (pending=10%, shipped=40%, delivered=40%, returned=10%)
- Category: electronics, clothing, home, etc.
- Seasonal date patterns for orders
- Unit tests verify distributions

## Related Constitution Sections
- **Verticalized > Generic (Principle I)**
- **Realism**: Real products and pricing
- **TDD Required (Development Practice 1)**
