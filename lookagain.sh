find -name "*.sh" -exec basename {} .sh \; | sed 's/\.sh$//' | sort -r
