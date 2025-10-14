# Feature Specification Prompt: F015 - Progress Tracking & Performance Optimization

## Feature Metadata
- **Feature ID**: F015
- **Name**: Progress Tracking & Performance Optimization
- **Category**: Data Generation
- **Phase**: Week 6
- **Priority**: P1 (Should-have)
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F013 (Data generation engine)

## Constitutional Alignment
- ✅ **Speed > Features**: <30s seeding time enforced
- ✅ **Developer-First Design**: Transparent progress, clear feedback
- ✅ **Performance Gates**: 100+ records/second

## User Story
**US-MVP-009**: "As a developer seeding a database, I want to see real-time progress with estimated completion time so I know the operation is working and how long it will take."

## Problem Statement
Without progress indication, users don't know if seeding is working or hung. Silent operations are bad UX. Need:
- Real-time progress bars
- Records/second rate
- Estimated time remaining
- Performance optimization to meet <30s target

## Solution Overview
Integrate progressbar library, add progress tracking to engine, optimize for >100 records/sec, display progress during generation.

## Acceptance Criteria
- Progress bar showing: Table name, Records (X/Y), %, ETA
- Performance: >100 records/sec
- Benchmark tests verify performance
- No silent operations

## Related Constitution Sections
- **Transparent Progress (UX Principle 3)**
- **Performance Gates (Technical Constraint 1)**
- **Speed > Features (Principle II)**
