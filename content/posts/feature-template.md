---
title: Feature Template
lastmod: "2024-07-25T20:55:30.003Z"
date: "2023-07-19T13:16:15.942Z"
---

```gherkin
Feature: NAME

  DESCRIPTION

  Background: NAME

  DESCRIPTION

  Given STEP_DESCRIPTION
  And STEP_DESCRIPTION

  Scenario Outline: NAME

  DESCRIPTION

  Given STEP_DESCRIPTION
  And STEP_DESCRIPTION "<input>"
  When STEP_DESCRIPTION
  But STEP_DESCRIPTION
  Then STEP_DESCRIPTION "<output>"
  And STEP_DESCRIPTION

    Example:
    | input  | output |
    | value  | result |
    | valve  | result |
```
