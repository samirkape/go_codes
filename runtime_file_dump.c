/************************************************************/
// runtime debug 

void cff(float *src_, int ndst_, const char *fname, char ipf)
{
  FILE *fp = fopen(fname, "w");
  char fsp[4] = {'%','f'};
  fsp[2] = ipf;
  int sz = sizeof(float) * ndst_;
  for (int i = 0; i < ndst_; i++)
    fprintf(fp, fsp, src_[i]);
  fflush(fp);
}
void cfd(double *src_, int ndst_, const char *fname, char ipf)
{
  char fsp[4] = {'%','f'};
  fsp[2] = ipf;
  FILE *fp = fopen(fname, "w");
  int sz = sizeof(float) * ndst_;
  for (int i = 0; i < ndst_; i++)
    fprintf(fp, fsp, src_[i]);
  fflush(fp);
}

void ps(double *src_, int ndst_)
{
  fprintf(stderr, "\n");
  for (int i = 0; i < ndst_; i++)
    fprintf(stderr, "%lf,", src_[i]);
  fflush(stderr);
}

void psf(float *src_, int ndst_)
{
  fprintf(stderr, "\n");
  for (int i = 0; i < ndst_; i++)
    fprintf(stderr, "%f,", src_[i]);
  fflush(stderr);
}
/***************************************************************/