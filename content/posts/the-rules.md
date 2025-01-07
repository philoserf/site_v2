---
title: The Rules
lastmod: "2025-01-07T01:18:18.077Z"
date: "2025-01-07T01:18:18.077Z"
---

## Keep it simple or Make it when you need it

- If it has a name, do not give it another name.
- Do not give it another name if you have given it a name.
- Start with one file until you need another file.
- Split content to a new file when its weight warrants a new file.
- Start with one folder until you need another folder.
- Create a subfolder when it is needed.

### Why These Matter

1. **Naming Consistency**\
   The first two rules emphasize the importance of avoiding aliases and maintaining a single source of truth. When multiple names exist for the same entity, it creates cognitive overhead and potential confusion. For example:

   - Bad: Having both `UserManager` and `UserHandler` classes that do the same thing
   - Good: Consistently using `UserManager` throughout the codebase

2. **Progressive Organization**\
   The file and folder rules promote organic growth of project structure. Start simple and only add complexity when justified by actual needs:
   - Bad: Creating an elaborate folder structure upfront with /src/components/ui/atoms/buttons/
   - Good: Starting with /components/ and only creating subfolders when you have enough related components to warrant organization

## Surrender or Have a reason

- Adapt to the default configuration.
- After a sincere effort to adapt fails, change the default configuration.
- Review your configuration for cruft from time to time.

### Why These Matter

1. **Default First**\
   This principle acknowledges that default configurations are often well-thought-out and battle-tested. Deviating from defaults should require justification:

   - Bad: Immediately customizing ESLint rules without understanding why the defaults exist
   - Good: Using default ESLint configuration, identifying specific pain points, then carefully modifying only necessary rules

2. **Mindful Customization**\
   When changes are needed, they should be deliberate and documented:

   - Bad: Copying configuration snippets from Stack Overflow without understanding them
   - Good: Documenting why specific configuration changes were necessary and what problems they solve

3. **Regular Maintenance**\
   Configuration debt is still technical debt. Regular reviews prevent accumulation of unnecessary customizations:
   - Bad: Keeping old configuration flags for deprecated features
   - Good: Periodically reviewing configuration files and removing unused settings

## Related Principles

### YAGNI (You Aren't Gonna Need It)

This principle aligns perfectly with "Make it when you need it." It suggests that you shouldn't add functionality until it's necessary:

- Bad: Building a complex plugin system because "we might need it later"
- Good: Writing simple, focused code that solves current problems

### DRY (Don't Repeat Yourself) with Caution

While avoiding duplication is important, it shouldn't override clarity or simplicity:

- Bad: Creating complex abstractions to eliminate any code duplication
- Good: Accepting some duplication if it makes the code more straightforward and maintainable

## Practical Application

1. **Project Structure**

   ```plain
   /project
   ├── index.js           # Start with one file
   ├── users.js           # Add when user logic grows
   └── /components        # Add folder when component count warrants it
   ```

2. **Configuration Evolution**

   ```javascript
   // Stage 1: Use defaults
   module.exports = {
     extends: "eslint:recommended",
   };

   // Stage 2: Add justified customizations
   module.exports = {
     extends: "eslint:recommended",
     rules: {
       // Added due to project-specific naming convention
       camelcase: ["error", { properties: "never" }],
     },
   };
   ```

## Todo

- \[x] Add _why_ & provide examples
- \[ ] Add more rules by others that align with these principles
- \[ ] Add section on when to break these rules
- \[ ] Include real-world case studies
