# Limbo

♟️ Limbo — online chess game.

# Todo for the future

- pkg/chess: Rewrite arch: separate packages, make more interfaces. Add more validations to the public code. Make more private functions. Make private unsafe functions without errors. Do not forget to rename error messages.
- pkg/chess: Rename function name: OccupiedByColor -> CheckSquareOccupiedByColor. Add argument names to function name where needed.
- pkg/chess: Remove **Raw** suffixes, add more docs about what function does (step by step).
- pkg/chess: Use SRP, for example, for FEN parsing.
- pkg/chess: Add the ability to use engine without models and models with custom engine.
- pkg/chess: Cache moves generation in engine and other long operations?
- pkg/chess: Benchmarks, speed, reduce memory usage.
- pkg/chess: Reduce allocations count.
- pkg/chess: Cache some calculations in fields.
- pkg/chess: Make enumerations more stable so as not to be afraid of unknown enumeration values.
