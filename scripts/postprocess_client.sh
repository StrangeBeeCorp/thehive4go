# Replace "var XXX int32 = now()" with proper Go code
sed -i 's/var \([a-zA-Z]*\) int32 = now()/var \1 int32 = int32(time.Now().Unix() * 1000)/' "$1"

# Replace "var XXX int64 = now()" with proper Go code (for datetime_ms fields converted to int64)
sed -i 's/var \([a-zA-Z]*\) int64 = now()/var \1 int64 = time.Now().Unix() * 1000/' "$1"

# Make sure time package is imported if we added time.Now()
if grep -q 'time.Now()' "$1"; then
    # Check if time import already exists
    if ! grep -q '^import "time"' "$1" && ! grep -q '"time"' "$1"; then
        # Add time import after package declaration
        sed -i '/^package /a\
import "time"' "$1"
    fi
fi
