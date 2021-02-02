
#include<stdio.h>
#include<stdlib.h>
#include<string.h>


double ** malloc2d( int rows, int cols  )
{
    double ** out_buf = calloc(rows , sizeof(double *)); 
        
    if ( out_buf == NULL )
    {
        fprintf(stderr, "Error in malloc\n");
        exit(-1);
    }

    for (int i = 0; i < rows; i++)
    {
        out_buf[i] = calloc(cols , sizeof(double));
        if ( out_buf[i] == NULL )
        {
            fprintf(stderr, "Error in malloc\n");
            exit(-1);
        }
    }
    return out_buf;
}


double ** read_csv( char * filename, int * out_rows, int * out_col )
{
    /***************** Variable Initializations ******************/
    FILE *input_csv = NULL;
    double ** out_featbuf = NULL;
    char *token;                         //store the resulted string output from strtok
    input_csv = fopen(filename,"r");
    char lines[500] = {0};                  //to fetch input line-by-line from i/p csv
    char tmp_lines[500] = {0};
    char cp_tmp[500] = {0};
    char tempstr[500] = {0};               //to preserve original string after using strtok
    char * search = ",\n";
    int column = 0;
    double first_token = 0.0;
    int len = 0;
    int rows = 0;
    int cols = 0;
    int num_chars_in_row = 0;
    
    // below code calculates number of rows 
    // without affecting input_csv pointer position

    if (input_csv)
    {
        fseek(input_csv, 0L, SEEK_END);
        len = ftell(input_csv);
        
        rewind(input_csv);        
        if (len > 0)    // get chars 
        {
            for (num_chars_in_row = 0; (fgetc(input_csv)) != 10; )  // 10 = ascii value of '\n'
            {
                num_chars_in_row++;
            }
        }
        
        rewind(input_csv);        
        if (len > 0)    // get cols
        {
            for (int i = 0; (fgets(tmp_lines, num_chars_in_row, input_csv)) != NULL; i++)
            {
                token = strtok( tmp_lines, "\n" );
                while(token!=NULL)
                {
                    token = strtok(NULL,"\n");
                    cols++;
                }
            }
        }
        
        rewind(input_csv);
        if (len > 0)    // get rows
        {
            if ( fgets(tmp_lines, num_chars_in_row, input_csv) != NULL )
            {
                strcpy( cp_tmp, tmp_lines );
                token = strtok( tmp_lines, "," );
                while( token!=NULL && strcmp(token, cp_tmp) )
                {
                    token = strtok( NULL, "," );
                    rows++;
                }
            }
        }
    }
    *out_rows = rows;   // number of features selected
    *out_col = cols;

    out_featbuf = malloc2d( rows, cols );

    if (input_csv)
    {
        rewind(input_csv);
        fseek(input_csv, 0L, SEEK_END);
        len = ftell(input_csv);
        rewind(input_csv);

        if (len > 0)
        {
            for (int i = 0; (fgets(lines, num_chars_in_row, input_csv)) != NULL; i++)
            {
                strcpy(tempstr, lines);
                token = strtok(tempstr, search);     //strtok first call

                first_token = atof(token);              //store output of first call 
                out_featbuf[0][column] = first_token;
                 while (token != NULL) 
                 { 
                    for( int i = 1; i < rows; i++ )
                    {
                        token = strtok(NULL, search);   
                        if(!token) break;    //rest calls of strtok till EOF   
                        out_featbuf[i][column] = atof(token);
                    }
                    if(!token)break;
                    column++;
                 }
            }
        }
    }
    #if DEBUG_FLAG 
            FILE* out;
            out = fopen("./write.csv","w");
            for(int j=0;j<rows;j++)
            {
                for(int i=0;i<2;i++)
                {
                    printf("%d - > %lf\t",j,out_featbuf[i][j]);
                }
                printf("\n");
            }
    #endif

    /************** memory release and input_csv closing *************/
    fclose(input_csv);
    return out_featbuf;
    /*******************************************************************/
}