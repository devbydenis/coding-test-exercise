var isValid = function(s) {
    const stack = []; // Stack untuk menyimpan kurung buka

    for (let i = 0; i < s.length; i++) { // Loop setiap karakter
        const char = s[i];

        if (char === '(' || char === '{' || char === '[') { // Jika ketemu kurung buka, simpan ke stack
            stack.push(char);
        } else { // Jika ketemu kurung tutup

            if (stack.length === 0) {   // Jika stack kosong tapi ada kurung tutup
                return false; // artinya tidak ada pasangan -> INVALID
            }

            const top = stack.pop();  // Ambil elemen terakhir dari stack (LIFO)

            // Cek apakah jenis kurung cocok
            if (char === ')' && top !== '(') return false;
            if (char === '}' && top !== '{') return false;
            if (char === ']' && top !== '[') return false;
        }
    }

    // Jika masih ada kurung buka yang belum ditutup
    return stack.length === 0;
};
