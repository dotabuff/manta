#include <stdint.h>

#define LZSS_LOOKSHIFT 4

int32_t lzss_uncompress(char *pInput, char *pOutput, int32_t size) {
    int32_t totalBytes = 0;
    int32_t cmdByte = 0;
    int32_t getCmdByte = 0;

    // ignore the header, 32 bit id, 32 bit size
    pInput += 8;

    for ( ;; ) {
        if ( !getCmdByte ) {
            cmdByte = *pInput++;
        }

        getCmdByte = ( getCmdByte + 1 ) & 0x07;

        if ( cmdByte & 0x01 ) {
            int32_t position = *pInput++ << LZSS_LOOKSHIFT;
            position |= ( *pInput >> LZSS_LOOKSHIFT );
            int32_t count = ( *pInput++ & 0x0F ) + 1;
            if ( count == 1 )
                break;

            char *pSource = pOutput - position - 1;
            for ( uint32_t i = 0; i < count; ++i ) {
                *pOutput++ = *pSource++;
            }

            totalBytes += count;
        } else {
            *pOutput++ = *pInput++;
            totalBytes++;
        }

        cmdByte = cmdByte >> 1;
    }

    if ( totalBytes != size ) {
        // unexpected failure
        // std::cout << "Total != actual: " << totalBytes << "," << size << std::endl;
        return totalBytes;
    } else {
        // std::cout << "Uncompressed" << std::endl;
    }

    return totalBytes;
}
